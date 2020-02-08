package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Josephnp732/ABDI/store"

	validator "github.com/xeipuuv/gojsonschema"
)

// Store JSON into the Redis DB
func createPlan(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	// Validate against JSON schema
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	httpLoad := validator.NewStringLoader(string(body))
	schema := validator.NewReferenceLoader("file://./schema.json")
	result, err := validator.Validate(schema, httpLoad)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		fmt.Println(err.Error())
	}

	if result.Valid() {
		http.Error(w, http.StatusText(201), 201)
		w.Write([]byte("JSON is valid. "))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("The document is not valid. see errors :\n"))
		for _, desc := range result.Errors() {
			w.Write([]byte((fmt.Sprintf("- %s\n", desc))))
		}
		return
	}

	// Unmarshall JSON into structure
	var plan store.Plan
	err = json.Unmarshal(body, &plan)
	if err != nil {
		fmt.Println("Unable to unmarshall JSON")
	}

	// Composite key stored as objectID_objectType
	key := fmt.Sprintf("%s_%s", plan.ObjectID, plan.ObjectType)

	// Store key Value pair in DB
	err = redisStore.CreateEntry(key, string(body))
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		fmt.Println(fmt.Sprintf("Redis Connection Error %s", err.Error()))
		return
	}
	w.Write([]byte("Successfully stored Key Value pair in DB"))
	return
}

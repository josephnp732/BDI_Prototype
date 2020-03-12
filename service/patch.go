package service

import (
	"fmt"
	"io/ioutil"
	"net/http"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
)

func patchPlan(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	// Authorize JWT
	AuthorizeUser(w, r)

	planID := chi.URLParam(r, "planID")
	planType := chi.URLParam(r, "planType")
	if planID == "" || planType == "" {
		w.Write([]byte("Please enter valid parameters"))
	}

	key := fmt.Sprintf("%s_%s", planID, planType)

	val, err := redisStore.RetrieveEntry(key)
	if err == redis.Nil {
		http.Error(w, http.StatusText(400), 400)
		w.Write([]byte(fmt.Sprintf("Value not found for key %s . Please check the PlanID", key)))
	} else if err != nil {
		http.Error(w, http.StatusText(500), 500)
		fmt.Printf("Redis Error: %s", err.Error())
	}

	// json object from redis
	original := []byte(val)

	// read the body
	target, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Merge the JSON body
	patch, err := jsonpatch.CreateMergePatch(original, target)
	if err != nil {
		panic(err)
	}

	// patch, err = jsonpatch.CreateMergePatch(original, patch)
	// if err != nil {
	// 	panic(err)
	// }

	// Store key Value pair in DB
	// err = redisStore.CreateEntry(key, string(patch))
	// if err != nil {
	// 	http.Error(w, http.StatusText(500), 500)
	// 	fmt.Println(fmt.Sprintf("Redis Connection Error %s", err.Error()))
	// 	return
	// }

	w.Write(patch)
}

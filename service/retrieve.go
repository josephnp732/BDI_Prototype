package service

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/amalfra/etag"

	"github.com/go-chi/chi"
	redis "github.com/go-redis/redis/v7"
)

func retrievePlan(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	planID := chi.URLParam(r, "planID")
	planType := chi.URLParam(r, "planType")
	if planID == "" || planType == "" {
		w.Write([]byte("Please enter valid parameters"))
	}

	key := fmt.Sprintf("%s_%s", planID, planType)

	// generate e-Tag
	eTagKey := etag.Generate(key, false)

	w.Header().Set("Etag", eTagKey)

	if match := r.Header.Get("If-None-Match"); match != "" {
		if strings.Contains(match, eTagKey) {
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}

	val, err := redisStore.RetrieveEntry(key)
	if err == redis.Nil {
		http.Error(w, http.StatusText(400), 400)
		w.Write([]byte(fmt.Sprintf("Value not found for key %s . Please check the PlanID", key)))
	} else if err != nil {
		http.Error(w, http.StatusText(500), 500)
		fmt.Printf("Redis Error: %s", err.Error())
	}
	w.Write([]byte(val))
}

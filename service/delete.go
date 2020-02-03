package service

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func deletePlan(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	planID := chi.URLParam(r, "planID")
	planType := chi.URLParam(r, "planType")
	if planID == "" || planType == "" {
		w.Write([]byte("Please enter valid parameters"))
	}

	key := fmt.Sprintf("%s_%s", planID, planType)

	resInt, err := redisStore.DeleteEntry(key)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Redis Connectipn Error: %s \n", err.Error())))
		http.Error(w, http.StatusText(500), 500)
	}

	if resInt > 0 {
		w.Write([]byte(fmt.Sprintf("Plan: %s with key %s successfully deleted", planType, planID)))
	} else {
		w.Write([]byte("Plan not found. Please check the parameters"))
	}

}

package service

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func deletePlan(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	key := chi.URLParam(r, "planID")
	if key == "" {
		w.Write([]byte("Please enter a vaid PlanID"))
	}

	if redisStore.DeleteEntry(key) > 0 {
		w.Write([]byte(fmt.Sprintf("Plan with key %s successfully deleted", key)))
	} else {
		w.Write([]byte("Plan not found. Please check the PlanID"))
	}

}

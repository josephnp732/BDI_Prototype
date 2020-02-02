package service

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	redis "github.com/go-redis/redis/v7"
)

func retrievePlan(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	key := chi.URLParam(r, "planID")
	if key == "" {
		w.Write([]byte("Please enter a vaid PlanID"))
	}

	val, err := redisStore.RetrieveEntry(key)
	if err == redis.Nil {
		w.Write([]byte(fmt.Sprintf("Value not found for key %s . Please check the PlanID", key)))
	} else if err != nil {
		fmt.Printf("Redis Error: %s", err.Error())
	}
	w.Write([]byte(val))
}

package service

import (
	"fmt"
	"net/http"

	"github.com/Josephnp732/ABDI/store"

	"github.com/go-chi/chi"
)

var redisStore = store.NewStore()

// MainHandler routes the requests to specific http functions
func MainHandler() {
	r := chi.NewRouter()
	fmt.Println("Application running on port :3000")
	r.Post("/plan", createPlan)
	r.Get("/ping", healthCheck)
	r.Delete("/delete/{planID}", deletePlan)
	r.Get("/retrieve/{planID}", retrievePlan)
	http.ListenAndServe(":3000", r)
}

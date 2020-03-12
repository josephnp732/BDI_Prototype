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

	// healthCheck
	r.Get("/ping", healthCheck)

	// get JWT Token
	r.Get("/token", token)

	r.Post("/plan", createPlan)
	r.Delete("/delete/{planID}/type/{planType}", deletePlan)
	r.Get("/retrieve/{planID}/type/{planType}", retrievePlan)
	r.Patch("/plan/{planID}/type/{planType}", patchPlan)
	http.ListenAndServe(":3000", r)
}

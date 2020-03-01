package service

import (
	"fmt"
	"net/http"

	"github.com/Josephnp732/ABDI/security"
)

// Health Check Handler
func token(w http.ResponseWriter, r *http.Request) {

	tokenString, err := security.GenerateJWT()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf(tokenString)))
	return
}

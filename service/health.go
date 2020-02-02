package service

import (
	"net/http"
)

// Health Check Handler
func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
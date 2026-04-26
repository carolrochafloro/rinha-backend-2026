package httpapi

import "net/http"

func NewRouter(handler *Handler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /ready", handler.Ready)
	mux.HandleFunc("POST /fraud-score", handler.FraudScore)

	return mux
}

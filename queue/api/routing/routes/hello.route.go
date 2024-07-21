package routes

import (
	"encoding/json"
	"net/http"
)

func SetHelloRoute(mux *http.ServeMux) {
	mux.HandleFunc("GET /hello/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("ping")
	})
}

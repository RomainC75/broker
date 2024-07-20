package routing

import (
	"fmt"
	"net/http"
)

func SetSocketRoute(mux *http.ServeMux) {
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Setting up the server!")
	})
}

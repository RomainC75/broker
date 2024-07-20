package api

import (
	"fmt"
	"net/http"
	"queue/config"
)

func Serve() {

	config.SetEnv()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Setting up the server!")
	})
	http.ListenAndServe(":8080", nil)
}

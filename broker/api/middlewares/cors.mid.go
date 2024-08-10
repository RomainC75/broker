package middlewares

import (
	"fmt"
	"net/http"
)

func CORSMiddleware(next http.HandlerFunc) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		// !
		if r.Method == "OPTIONS" {
			// status : 204
			w.WriteHeader(http.StatusNoContent)
			w.Write([]byte{})
			fmt.Println("--> send ")
			return
		}

		next.ServeHTTP(w, r)
	})
}

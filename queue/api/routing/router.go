package routing

import "net/http"

func SetRouter(mux *http.ServeMux) {
	SetSocketRoute(mux)
}

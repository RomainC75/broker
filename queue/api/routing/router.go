package routing

import (
	"net/http"
	"queue/api/routing/routes"
)

func SetRouter(mux *http.ServeMux) {
	routes.SetSocketRoute(mux)
	routes.SetHelloRoute(mux)
}

package routing

import (
	"broker/api/routing/routes"
	"net/http"
)

func SetRouter(mux *http.ServeMux) {
	routes.SetSocketRoute(mux)
	routes.SetHelloRoute(mux)
}

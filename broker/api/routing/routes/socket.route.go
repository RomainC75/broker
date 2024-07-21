package routes

import (
	"broker/api/controllers"
	"net/http"

	"golang.org/x/net/websocket"
)

func SetSocketRoute(mux *http.ServeMux) {
	socketServer := controllers.NewSocketCtrl()
	mux.Handle("/ws", websocket.Handler(socketServer.HandleWorkTest))
}

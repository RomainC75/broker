package routes

import (
	"net/http"
	"queue/api/controllers"

	"golang.org/x/net/websocket"
)

func SetSocketRoute(mux *http.ServeMux) {
	socketServer := controllers.NewSocketCtrl()
	mux.Handle("/ws", websocket.Handler(socketServer.HandleWorkTest))
}

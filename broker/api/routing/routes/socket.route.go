package routes

import (
	"broker/api/controllers"
	"context"
	"net/http"

	"golang.org/x/net/websocket"
)

func SetSocketRoute(mux *http.ServeMux) {
	ctx := context.Background()
	socketServer := controllers.NewSocketCtrl(ctx)
	mux.Handle("/ws", websocket.Handler(socketServer.HandleWorkTest))
}

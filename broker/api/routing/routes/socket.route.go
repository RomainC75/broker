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
	mux.HandleFunc("/socket/ticket", socketServer.HandleTicket)
	mux.Handle("/socket/ws", websocket.Handler(socketServer.HandleBroker))
	mux.Handle("/socket/reader", websocket.Handler(socketServer.HandleWatch))

}

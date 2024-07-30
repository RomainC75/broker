package controllers

import (
	"broker/broker"
	"context"
	"fmt"

	"golang.org/x/net/websocket"
)

type SocketCtrl struct {
	Broker *broker.Broker
}

func NewSocketCtrl(ctx context.Context) *SocketCtrl {
	b := broker.NewBroker()
	b.LaunchLoop(ctx)
	return &SocketCtrl{
		Broker: b,
	}
}

func (socketCtrl *SocketCtrl) HandleBroker(conn *websocket.Conn) {
	socketCtrl.Broker.AddClient(conn)
	fmt.Println("client added")
	// fmt.Println("=> <", req)
	// ctrl_utils.SendJsonResponse(w, http.StatusCreated, ctrl_utils.CtrlResponse{"message": "created"})
}

func (socketCtrl *SocketCtrl) HandleWatch(conn *websocket.Conn) {
	socketCtrl.Broker.AddWatcher(conn)
	fmt.Println("client added")
	// fmt.Println("=> <", req)
	// ctrl_utils.SendJsonResponse(w, http.StatusCreated, ctrl_utils.CtrlResponse{"message": "created"})
}

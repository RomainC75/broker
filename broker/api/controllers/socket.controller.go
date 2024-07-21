package controllers

import (
	"broker/broker"
	"fmt"

	"golang.org/x/net/websocket"
)

type SocketCtrl struct {
	Broker *broker.Broker
}

func NewSocketCtrl() *SocketCtrl {
	b := broker.NewBroker()
	return &SocketCtrl{
		Broker: b,
	}
}

func (socketCtrl *SocketCtrl) HandleWorkTest(conn *websocket.Conn) {
	socketCtrl.Broker.AddClient(conn)
	fmt.Println("client added")
	// fmt.Println("=> <", req)
	// ctrl_utils.SendJsonResponse(w, http.StatusCreated, ctrl_utils.CtrlResponse{"message": "created"})

}

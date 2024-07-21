package controllers

import (
	"fmt"
	"queue/broker"

	"golang.org/x/net/websocket"
)

type SocketCtrl struct {
	Broker *broker.Broker
}

func NewSocketCtrl() *SocketCtrl {
	return &SocketCtrl{
		Broker: broker.GetBroker(),
	}
}

func (socketCtrl *SocketCtrl) HandleWorkTest(conn *websocket.Conn) {
	fmt.Println("sdf")
	// fmt.Println("=> <", req)
	// ctrl_utils.SendJsonResponse(w, http.StatusCreated, ctrl_utils.CtrlResponse{"message": "created"})
}

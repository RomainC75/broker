package controllers

import (
	"broker/api/services"
	"broker/broker"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"shared/utils"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
)

type SocketCtrl struct {
	Broker     *broker.Broker
	SsoService services.SsoServiceInterface
}

func NewSocketCtrl(ctx context.Context) *SocketCtrl {
	b := broker.NewBroker()
	b.LaunchLoop(ctx)
	b.LaunchWatcherLoop(ctx)
	return &SocketCtrl{
		Broker:     b,
		SsoService: services.NewSsoService(),
	}
}

func (SocketCtrl *SocketCtrl) HandleTicket(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	fmt.Println("=>", auth)

	authorizationParts := strings.Split(auth, " ")
	if len(authorizationParts) != 2 {
		json.NewEncoder(w).Encode("bearer error ")
	}

	userClaims, err := SocketCtrl.SsoService.ExtractTokenClaims(authorizationParts[1])
	if err != nil {
		logrus.Error(err.Error())
	}

	utils.PrettyDisplay("RESULT : ", userClaims)
	json.NewEncoder(w).Encode("ping")

}

func (socketCtrl *SocketCtrl) HandleBroker(conn *websocket.Conn) {
	socketCtrl.Broker.AddClient(conn)
	fmt.Println("client added")
	// fmt.Println("=> <", req)
	// ctrl_utils.SendJsonResponse(w, http.StatusCreated, ctrl_utils.CtrlResponse{"message": "created"})
}

func (socketCtrl *SocketCtrl) HandleWatch(conn *websocket.Conn) {
	time.Sleep(time.Second)
	socketCtrl.Broker.AddWatcher(conn)
	fmt.Println("client added")
	// fmt.Println("=> <", req)
	// ctrl_utils.SendJsonResponse(w, http.StatusCreated, ctrl_utils.CtrlResponse{"message": "created"})
}

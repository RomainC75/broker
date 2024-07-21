package controllers

import (
	"context"
	"net/http"
	"queue/broker"

	"github.com/hibiken/asynq"
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
	workCode := r.PathValue("work_code")
	userId := r.Context().Value("user_id").(int32)

	ctx := context.Background()
	distributor := events.Get()
	opts := []asynq.Option{
		asynq.MaxRetry(10),
		// asynq.ProcessIn(time.Second),
		// + send to other queue :-)
		asynq.Queue(string(events.CriticalQueueReq)),
	}

	scenarioRequest := work_dto.PortTestScenarioRequest{
		UserId:    userId,
		IPRange:   portTestScenario.PortTestScenario.IPRange,
		PortRange: portTestScenario.PortTestScenario.PortRange,
	}

	err := distributor.DistributeTaskSendWork(
		ctx,
		&scenarioRequest,
		events.PortScanner,
		opts...,
	)
	if err != nil {
		ctrl_utils.SendJsonResponse(w, http.StatusCreated, ctrl_utils.CtrlResponse{"message": "created"})
	}

	// fmt.Println("=> <", req)
	ctrl_utils.SendJsonResponse(w, http.StatusCreated, ctrl_utils.CtrlResponse{"message": "created"})
}

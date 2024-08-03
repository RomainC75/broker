package main

import (
	"context"
	"fmt"
	"net/url"
	"producer/dummy"
	message_broker "shared/broker"
	"shared/config"
	"time"

	"sync"

	"github.com/google/uuid"
)

var (
	origin = "http://localhost"
)

func main() {
	time.Sleep(time.Second)

	// * config
	producerName := uuid.New()
	config.SetEnv()
	conf := config.Getenv()
	topic := conf.BrokerTopic
	u := url.URL{Scheme: "ws", Host: fmt.Sprintf("%s:%s", conf.BrokerHost, conf.BrokerPort), Path: "/ws"}

	var wg sync.WaitGroup
	wg.Add(1)

	// * connection * //
	mb_conn := message_broker.NewConn(u, origin)

	// * produce * //
	ctx := context.Background()
	dummy.GoLoopProducer(producerName.String(), topic, mb_conn.Produce, time.Second*2, ctx)
	wg.Add(1)

	wg.Wait()
}

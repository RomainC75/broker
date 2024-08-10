package main

import (
	"context"
	"producer/binance"
	"shared/config"
	"time"

	"sync"
)

var (
	origin = "http://localhost"
)

func main() {
	time.Sleep(time.Second)

	// * config
	config.SetEnv()
	conf := config.Getenv()
	// u := url.URL{Scheme: "ws", Host: fmt.Sprintf("%s:%s", conf.BrokerHost, conf.BrokerPort), Path: "/socket/ws"}

	var wg sync.WaitGroup
	wg.Add(1)

	ctx := context.Background()

	bin := binance.NewConn()
	bin.GoListen(conf.BrokerTopic, ctx)

	// * broker * //
	// producerName := uuid.New()
	// topic := conf.BrokerTopic
	// mb_conn := message_broker.NewConn(u, origin)
	// dummy.GoLoopProducer(producerName.String(), topic, mb_conn.Produce, time.Second*2, ctx)

	wg.Add(1)

	wg.Wait()
}

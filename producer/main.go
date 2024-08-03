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
)

var (
	origin = "http://localhost"
)

func main() {
	time.Sleep(time.Second)

	config.SetEnv()
	conf := config.Getenv()
	u := url.URL{Scheme: "ws", Host: fmt.Sprintf("%s:%s", conf.BrokerHost, conf.BrokerPort), Path: "/ws"}

	var wg sync.WaitGroup
	wg.Add(1)

	mb_conn := message_broker.NewConn(u, origin)
	topic := conf.BrokerTopic

	// mb_conn.Produce(topic, []byte("hello1"))
	// time.Sleep(time.Second)
	// mb_conn.Produce(topic, []byte("hello2"))

	ctx := context.Background()
	dummy.GoLoopProducer(topic, mb_conn.Produce, time.Second*2, ctx)
	wg.Add(1)

	wg.Wait()
}

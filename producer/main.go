package main

import (
	"fmt"
	"net/url"
	message_broker "shared/broker"
	"shared/config"

	"sync"
)

var (
	origin = "http://localhost"
)

func main() {
	config.SetEnv()
	conf := config.Getenv()

	u := url.URL{Scheme: "ws", Host: fmt.Sprintf("%s:%s", conf.BrokerHost, conf.BrokerPort), Path: "/ws"}

	var wg sync.WaitGroup
	wg.Add(1)

	mb_conn := message_broker.NewConn(u, origin)
	topic := conf.BrokerTopic

	mb_conn.Produce(2, topic, []byte("hello"))

	wg.Wait()
}

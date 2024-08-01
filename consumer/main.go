package main

import (
	"fmt"
	"net/url"
	message_broker "shared/broker"
	"shared/config"
	"sync"
	"time"
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
	mb_conn.Subscribe(topic)

	mb_conn.GoHandleJobs(jobHandler)

	wg.Wait()
}

func jobHandler(message []byte) bool {
	fmt.Printf("handler got message ")
	return true
}

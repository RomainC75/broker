package main

import (
	"fmt"
	"net/url"
	"producer/conf"
	message_broker "shared/broker"
	"sync"
)

var (
	u      = url.URL{Scheme: "ws", Host: "localhost:3005", Path: "/ws"}
	origin = "http://localhost"
	topic  = "my_topic"
)

func main() {
	conf.LoadEnv()

	var wg sync.WaitGroup
	wg.Add(1)

	mb_conn := message_broker.NewConn(u, origin)
	mb_conn.Subscribe(topic)
	mb_conn.GoHandleJobs(jobHandler)

	wg.Wait()
}

func jobHandler(message []byte) bool {
	fmt.Printf("handler got message ")
	return true
}

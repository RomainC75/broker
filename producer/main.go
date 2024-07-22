package main

import (
	"net/url"
	"producer/binance"
	"producer/conf"
	message_broker "shared/broker"

	"sync"
)

var (
	u      = url.URL{Scheme: "ws", Host: "localhost:3005", Path: "/ws"}
	origin = "http://localhost"
)

func main() {
	conf.LoadEnv()

	var wg sync.WaitGroup
	wg.Add(1)

	mb_conn := message_broker.NewConn(u, origin)
	mb_conn.Produce(2, []byte("hello"))

	conn := binance.NewConn()
	conn.GoListen()

	wg.Wait()
}

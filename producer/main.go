package main

import (
	"producer/binance"
	"producer/conf"
	"producer/mb_broker"

	"sync"
)

func main() {
	conf.LoadEnv()

	var wg sync.WaitGroup
	wg.Add(1)

	mb_conn := mb_broker.NewConn()
	mb_conn.Produce(2, []byte("hello"))

	conn := binance.NewConn()
	conn.GoListen()

	wg.Wait()
}

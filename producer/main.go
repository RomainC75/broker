package main

import (
	"producer/conf"
	"producer/kafka"
	"producer/socket"
	"sync"
)

func main() {
	conf.LoadEnv()

	var wg sync.WaitGroup
	wg.Add(1)

	conn := socket.NewConn()
	conn.GoListen()

	kafka.NewConn()
	kafka.Produce(2, "hello")

	wg.Wait()
}

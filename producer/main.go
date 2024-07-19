package main

import (
	"producer/conf"
	"producer/kafka"
	"producer/socket"
	"sync"
)

func main() {
	conf.LoadEnv()

	kafka.NewConnection()

	var wg sync.WaitGroup
	wg.Add(1)

	conn := socket.NewConn()
	conn.GoListen()

	wg.Wait()
}

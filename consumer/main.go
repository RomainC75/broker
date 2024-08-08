package main

import (
	"fmt"
	"net/url"
	message_broker "shared/broker"
	"shared/config"
	db "shared/db/sqlc"
	"sync"
	"time"
)

var (
	origin = "http://localhost"
)

func main() {
	time.Sleep(time.Second * 3)

	config.SetEnv()
	db.Connect()

	conf := config.Getenv()

	u := url.URL{Scheme: "ws", Host: fmt.Sprintf("%s:%s", conf.BrokerHost, conf.BrokerPort), Path: "/ws"}

	var wg sync.WaitGroup
	wg.Add(1)

	mb_conn := message_broker.NewConn(u, origin)

	mb_conn.SubscribeTopics(conf.Tickers)

	mb_conn.GoHandleJobs(jobHandler)

	wg.Wait()
}

func jobHandler(message []byte) bool {
	time.Sleep(time.Second * 3)
	fmt.Println("--- GOT MESSAGE : ", string(message))
	return true
}

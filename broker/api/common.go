package api

import (
	"broker/api/routing"
	"broker/broker"
	"broker/config"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var mux *http.ServeMux

func Init() {
	mux = http.NewServeMux()
}

func GetRouter() *http.ServeMux {
	routing.SetRouter(mux)
	return mux
}

func Serve() {

	config.SetEnv()
	conf := config.Getenv()
	Init()
	mux = GetRouter()

	// mux.ListenAndServe(":8080", nil)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", conf.BrokerPort),
		Handler: mux,
	}

	go func() {
		fmt.Printf("====> listening to port : %s\n", conf.BrokerPort)
		http.ListenAndServe(fmt.Sprintf(":%s", conf.BrokerPort), mux)
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	broker := broker.GetBroker()
	broker.CloseEveryConnections()

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	<-ctx.Done()
	log.Println("Server exiting")
}

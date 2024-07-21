package kafka

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"golang.org/x/net/websocket"
)

var connection *Connection

type Connection struct {
	url    url.URL
	config *websocket.Config
	conn   *websocket.Conn
	ctx    context.Context
}

var (
	u = url.URL{Scheme: "ws", Host: "localhost:3005", Path: "/ws"}
)

func NewConn() *Connection {

	config, err := websocket.NewConfig(u.String(), "http://localhost")
	if err != nil {
		log.Fatal("error with config: ", err.Error())
	}
	ctx := context.Background()
	conn, err := config.DialContext(ctx)
	if err != nil {
		log.Fatal("error trying to dial: ", err.Error())
	}

	connection = &Connection{
		url:    u,
		config: config,
		conn:   conn,
		ctx:    ctx,
	}
	return connection
}

func Produce(i int, message string) {
	fmt.Println("producing : ", message)
	// to produce messages
	if connection == nil {
		fmt.Println("no wriiter")
	}
	_, err := connection.conn.Write([]byte(message))
	if err != nil {
		log.Fatal("failed to write messages:", err.Error())
	}
}

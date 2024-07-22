package mb_broker

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
	u      = url.URL{Scheme: "ws", Host: "localhost:3005", Path: "/ws"}
	origin = "http://localhost"
)

func GetConnection() *Connection {
	return connection
}

func NewConn() *Connection {

	config, err := websocket.NewConfig(u.String(), origin)
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

func (c *Connection) Produce(i int, message []byte) {
	fmt.Println("producing : ", message)
	// to produce messages
	if connection == nil {
		fmt.Println("no wriiter")
	}
	_, err := connection.conn.Write(message)
	if err != nil {
		log.Fatal("failed to write messages:", err.Error())
	}
}

package message_broker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"shared/broker_dto"

	"golang.org/x/net/websocket"
)

var connection *Connection

type Connection struct {
	url    url.URL
	config *websocket.Config
	conn   *websocket.Conn
	ctx    context.Context
}

func NewConn(u url.URL, origin string) *Connection {

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

func (c *Connection) Subscribe(topic string) {
	message := broker_dto.TopicSelection{
		Topic:   topic,
		Request: broker_dto.Subscribe,
	}
	b, err := json.Marshal(message)
	if err != nil {
		fmt.Println("=> ", err.Error())
		fmt.Println("imposible to marshall this message : ", message)
	}
	c.SendMessage(b)
}

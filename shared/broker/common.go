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

func GetConnection() *Connection {
	return connection
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
	message := broker_dto.Message{
		Topic:      topic,
		ActionCode: broker_dto.Subscribe,
	}
	b, err := json.Marshal(message)
	if err != nil {
		fmt.Println("=> ", err.Error())
		fmt.Println("imposible to marshall this message : ", message)
	}
	c.SendMessage(b)
}

func (c *Connection) Produce(i int, topic string, message []byte) {
	fmt.Println("producing : ", message)
	// to produce messages
	if connection == nil {
		fmt.Println("no wriiter")
	}
	mess := broker_dto.Message{
		Topic:      topic,
		ActionCode: broker_dto.SendMessage,
		Content:    message,
	}

	b, err := json.Marshal(mess)
	if err != nil {
		fmt.Println("error trying to marshall broker_dto", err.Error())
	}
	_, err = connection.conn.Write(b)
	if err != nil {
		log.Fatal("failed to write messages:", err.Error())
	}
}

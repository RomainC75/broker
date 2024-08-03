package message_broker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"shared/broker_dto"
	"shared/utils"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
)

var connection *Connection

type Connection struct {
	url    url.URL
	config *websocket.Config
	conn   *websocket.Conn
	ctx    context.Context
	busy   bool
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

func (c *Connection) Produce(topic string, message []byte) {
	fmt.Println("produce : ", message)
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

func (c *Connection) GoHandleJobs(handlerFn func([]byte) bool) {
	go func() {
		for {
			msg := make([]byte, 2048)
			_, err := c.conn.Read(msg)
			if err != nil {
				c.conn.Close()
				logrus.Error("error trying to read the conn")
				return
			}

			newMsg := utils.CleanByte(msg)
			var messageContent broker_dto.Message
			err = json.Unmarshal(newMsg, &messageContent)
			if err != nil {
				fmt.Println("error trying to unmarshal request : ", err.Error())
			}
			utils.PrettyDisplay("CONSUMER", messageContent)

			switch messageContent.ActionCode {
			// !! should be a separate go routine
			case broker_dto.Ping:
				c.SendPong()
			// !! shoud be a separate go routine
			case broker_dto.IsAvailable:
				c.SendIsAvailableInfo(true)
			case broker_dto.SendJob:
				c.SendIsAvailableInfo(false)
				c.SendAcceptJobMessage(messageContent.Topic, messageContent.Offset)
				handlerFn(messageContent.Content)
				c.SendIsAvailableInfo(true)
			}
		}
	}()
}

func (c Connection) SendIsAvailableInfo(isAvailable bool) {
	message, err := broker_dto.GetIsAvailableMessage(isAvailable)
	if err != nil {
		fmt.Println("sendIsAvailableInfo : cannot get isAvailable message")
	}
	b, err := json.Marshal(message)
	if err != nil {
		fmt.Println("marshall : isAvailable response not possible")
	}
	_, err = c.conn.Write(b)
	if err != nil {
		fmt.Println("response isAvailable not possible")
	}
}

func (c Connection) SendAcceptJobMessage(topic string, offset int) {
	message := broker_dto.Message{
		Topic:      topic,
		ActionCode: broker_dto.AcceptJob,
		Offset:     offset,
	}
	logrus.Warn("sned accept message ")
	utils.PrettyDisplay("SEND ACCEPT MEESSAGE ", message)
	b, err := json.Marshal(message)
	if err != nil {
		fmt.Println("marshall : isAvailable response not possible")
	}
	_, err = c.conn.Write(b)
	if err != nil {
		fmt.Println("accept job message sent !")
	}
}

func (c Connection) SendPong() {
	msg := broker_dto.Message{
		ActionCode: broker_dto.Pong,
	}
	b, err := json.Marshal(msg)
	if err != nil {
		fmt.Printf("marshall : pong not possible")
	}
	_, err = c.conn.Write(b)
	if err != nil {
		fmt.Printf("pong not possible")
	}
}

func SendMessage(c Connection, message broker_dto.Message) {

}

package message_broker

import (
	"fmt"
	"log"
)

func (c *Connection) SendMessage(message []byte) {
	fmt.Println("send message : ", message)
	// to produce messages
	if connection == nil {
		fmt.Println("no wriiter")
	}
	_, err := connection.conn.Write([]byte(message))
	if err != nil {
		log.Fatal("failed to write messages:", err.Error())
	}
}

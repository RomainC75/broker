package broker

import (
	"fmt"

	"golang.org/x/net/websocket"
)

func (b *Broker) Launch() {

}

func (b *Broker) AddClient(conn *websocket.Conn) {
	newClient := NewClient(conn)
	b.Clients[newClient] = true
	b.GoListenToClient(newClient)
}

func (b *Broker) GoListenToClient(client *Client) {
	go func() {
		for {
			msg := make([]byte, 2048)
			_, err := client.Conn.Read(msg)
			if err != nil {
				fmt.Println("error trying to read socket message")
			}
			fmt.Println("message : ", string(msg))
		}
	}()
}

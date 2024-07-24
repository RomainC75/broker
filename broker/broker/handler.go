package broker

import (
	"encoding/json"
	"fmt"
	"shared/broker_dto"
	"shared/utils"
	"sync"

	"golang.org/x/net/websocket"
)

func (b *Broker) Launch() {

}

func (b *Broker) AddClient(conn *websocket.Conn) {
	var wg sync.WaitGroup
	wg.Add(1)
	newClient := NewClient(conn)
	b.Clients[newClient] = true
	b.GoListenToClient(newClient, &wg)
	wg.Wait()
}

func (b *Broker) RemoveClient(client *Client) {
	b.Clients[client] = false
}

func (b *Broker) GoListenToClient(client *Client, wg *sync.WaitGroup) {
	go func() {
		defer client.Conn.Close()
		for {
			msg := make([]byte, 2048)
			_, err := client.Conn.Read(msg)
			newMsg := CleanByte(msg)
			if err != nil {
				fmt.Println("error trying to read socket message", err.Error())
				// remove client/ stop parent & current function
				b.RemoveClient(client)
				wg.Done()
				return
			} else {
				var messageContent broker_dto.Message
				err := json.Unmarshal(newMsg, &messageContent)
				if err != nil {
					fmt.Println("error trying to unmarshal request : ", err.Error())
				}
				utils.PrettyDisplay("request", messageContent)
				switch messageContent.ActionCode {
				case 0:
					fmt.Println("trying to unsubscribe")
				case 1:
					fmt.Printf("trying to subscribe to topic : %s\n", messageContent.Topic)
					if _, ok := b.Topics[messageContent.Topic]; ok {
						b.addClientToTopic(messageContent.Topic, client)
					} else {
						b.Topics[messageContent.Topic] = NewTopic(messageContent.Topic)
					}
				case 2:
					fmt.Println("tying to send message")
				}
			}
		}
	}()
}

func CleanByte(b []byte) []byte {
	position := len(b) - 1
	for b[position] == '\x00' {
		position--
	}
	newB := make([]byte, position+1)
	fmt.Println("===> position : ", position)
	fmt.Println("len b", len(b), b)
	copy(newB, b[:position+1])
	fmt.Println("len b", len(newB), newB)
	return newB
}

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
				fmt.Println("message : ", string(newMsg))
				utils.PrettyDisplay("request", messageContent)
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
	fmt.Println("len b", len(b))
	fmt.Println("len b", len(newB))
	copy(newB, b[:position])
	return newB
}

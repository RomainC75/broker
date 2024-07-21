package broker

import (
	"fmt"
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

func (b *Broker) GoListenToClient(client *Client, wg *sync.WaitGroup) {
	go func() {
		defer client.Conn.Close()
		for {
			msg := make([]byte, 2048)
			_, err := client.Conn.Read(msg)
			if err != nil {
				fmt.Println("error trying to read socket message", err.Error())
				wg.Done()
				return
			} else {
				fmt.Println("message : ", string(msg))
			}
		}
	}()
}

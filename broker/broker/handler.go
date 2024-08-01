package broker

import (
	"encoding/json"
	"fmt"
	"shared/broker_dto"
	"shared/utils"
	"sync"

	"golang.org/x/net/websocket"
)

func (b *Broker) AddClient(conn *websocket.Conn) {
	var wg sync.WaitGroup
	wg.Add(1)
	newClient := NewClient(conn)
	b.m.Lock()
	b.Clients[newClient] = true
	b.m.Unlock()
	b.GoListenToClient(newClient, &wg)
	wg.Wait()
}

func (b *Broker) RemoveClient(client *Client) {
	b.Clients[client] = false
}

// * =======  LOOP =======
func (b *Broker) GoListenToClient(client *Client, wg *sync.WaitGroup) {
	go func() {
		defer client.Conn.Close()
		for {
			msg := make([]byte, 2048)
			_, err := client.Conn.Read(msg)
			newMsg := utils.CleanByte(msg)
			fmt.Println("------------------------NEW MESSAAGE RECEIVED----------------------------")
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
				// utils.PrettyDisplay("request", messageContent)
				switch messageContent.ActionCode {
				case broker_dto.UnSubscribe:
					fmt.Println("trying to unsubscribe")
				case broker_dto.Subscribe:
					fmt.Printf("trying to subscribe to topic : %s\n", messageContent.Topic)
					isTopicExist := b.isTopicExists(messageContent.Topic)
					if !isTopicExist {
						b.Topics[messageContent.Topic] = NewTopic()
					}
					b.addClientToTopic(messageContent.Topic, client)
				case broker_dto.SendMessage:
					fmt.Println("tying to push message into queue")
					b.addMessage(messageContent)
				case broker_dto.IsAvailable:
					var isAvailableDto broker_dto.IsAvailableContent
					err := json.Unmarshal(messageContent.Content, &isAvailableDto)
					if err != nil {
						fmt.Println("error tryin to un marshal isAvailableDto")
					}
					client.SetIsAvailable(isAvailableDto.IsAvailable)
				case broker_dto.AcceptJob:
					b.SetJobToAccepted(messageContent.Topic, messageContent.Offset)
				}
			}

		}

	}()
}

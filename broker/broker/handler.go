package broker

import (
	"encoding/json"
	"fmt"
	"shared/broker_dto"
	"shared/utils"
	"sync"

	"github.com/sirupsen/logrus"
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
	b.m.Lock()
	b.Clients[client] = false
	b.m.Unlock()
}

// * =======  LOOP =======
func (b *Broker) GoListenToClient(client *Client, wg *sync.WaitGroup) {
	go func() {
		defer client.Conn.Close()
		for {
			msg := make([]byte, 2048)
			_, err := client.Conn.Read(msg)
			newMsg := utils.CleanByte(msg)

			if err != nil {
				logrus.Error("error trying to read socket message", err.Error())
				b.RemoveClient(client)
				wg.Done()
				return
			} else {
				var messageContent broker_dto.Message
				err := json.Unmarshal(newMsg, &messageContent)
				if err != nil {
					fmt.Println("error trying to unmarshal request : ", err.Error())
				}

				// *  CHOICE  * //

				switch messageContent.ActionCode {
				case broker_dto.UnSubscribe:
					b.removeClientFromTopic(messageContent.Topic, client)
				case broker_dto.Subscribe:
					isTopicExist := b.isTopicExists(messageContent.Topic)
					if !isTopicExist {
						b.Topics[messageContent.Topic] = NewTopic()
					}
					b.addClientToTopic(messageContent.Topic, client)
				case broker_dto.SendMessage:
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

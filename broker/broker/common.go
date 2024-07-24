package broker

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/net/websocket"
)

var broker *Broker

type PingInfo struct {
	LastPing   time.Time
	IsPingSent bool
	IsPong     bool
}

type Client struct {
	Conn        *websocket.Conn
	Ping        PingInfo
	Topics      []string
	IsAvailable bool
}

type Message struct {
	Key       []byte
	Value     []byte
	IsSent    bool
	IsHandled bool
}

type Topic struct {
	Content []Message
	// ConsumerCients []*Client
	ConsumerCients map[*Client]bool
}

type Broker struct {
	Clients map[*Client]bool
	Topics  map[string]Topic
}

func NewBroker() *Broker {
	broker = &Broker{
		Clients: map[*Client]bool{},
		Topics:  map[string]Topic{},
	}
	return broker
}

func (b *Broker) Launch(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				b.scanTopicsAndSend()
				time.Sleep(time.Millisecond * 100)
			}
		}
	}()
}

func GetBroker() *Broker {
	return broker
}

func (b *Broker) scanTopicsAndSend() {
	for _, topic := range b.Topics {
		fmt.Println("-> scanning : LAST CONTENT : ", topic.Content[len(topic.Content)-1])
	}
}

// Connection to topic
// send message key/value

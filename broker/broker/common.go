package broker

import (
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

func GetBroker() *Broker {
	return broker
}

// Connection to topic
// send message key/value

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
	Conn   *websocket.Conn
	Ping   PingInfo
	Topics []string
}

type Message struct {
	Key       []byte
	Value     []byte
	IsSent    bool
	IsHandled bool
}

type Topic struct {
	Name    string
	Content []Message
}

type Broker struct {
	Clients map[*Client]bool
	Topics  []Topic
}

func NewBroker() *Broker {
	broker = &Broker{
		Clients: map[*Client]bool{},
		Topics:  []Topic{},
	}
	return broker
}

func GetBroker() *Broker {
	return broker
}

// Connection to topic
// send message key/value

package broker

import (
	"time"

	"golang.org/x/net/websocket"
)

var broker *Broker

type Producer struct {
	Socket *websocket.Conn
	Ping   struct {
		LastPing time.Time
		IsPong   bool
	}
}

type Consumer struct {
	Socket *websocket.Conn
	Ping   struct {
		LastPing time.Time
		IsPong   bool
	}
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
	Producers []Producer
	Consumers []Consumer
	Topics    []Topic
}

func NewBroker() *Broker {
	broker = &Broker{}
	return broker
}

func GetBroker() *Broker {
	return broker
}

// Connection to topic
// send message key/value

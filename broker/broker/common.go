package broker

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/net/websocket"
)

var broker *Broker

type PingInfo struct {
	LastPing   time.Time
	IsPingSent bool
	IsPong     bool
	Retry      int
}

type Client struct {
	Conn        *websocket.Conn
	Ping        PingInfo
	Topics      []string
	IsAvailable bool
	m           *sync.Mutex
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
	ReaderIndex    int
}

type PingParameter struct {
	IntervalAfterPing time.Duration
	IntervalAfterPong time.Duration
	MaxRetry          int
}

type BrokerParameters struct {
	Ping PingParameter
}

type Broker struct {
	Clients    map[*Client]bool
	Topics     map[string]Topic
	Parameters BrokerParameters
}

func NewBroker() *Broker {
	broker = &Broker{
		Clients: map[*Client]bool{},
		Topics:  map[string]Topic{},
		Parameters: BrokerParameters{
			Ping: PingParameter{
				IntervalAfterPing: time.Second * 1,
				IntervalAfterPong: time.Second * 5,
				MaxRetry:          3,
			},
		},
	}
	return broker
}

func (b *Broker) CloseEveryConnections() {
	nb := 0
	for c := range b.Clients {
		c.Close()
		nb++
	}
	fmt.Printf("--> %d connections closed\n", nb)
}

func (b *Broker) Launch(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				b.scanTopicsAndSend()
				b.scanForPing()
				time.Sleep(time.Second)
			}
		}
	}()
}

func GetBroker() *Broker {
	return broker
}

func (b *Broker) scanTopicsAndSend() {
	fmt.Println("SCANNER")
	for topicName, topic := range b.Topics {
		fmt.Println("->", topicName, len(topic.Content))
		fmt.Println("-> listeners : ", len(topic.ConsumerCients))
		if len(topic.Content) > 0 {
			fmt.Println("-> scanning : LAST CONTENT : ", topic.Content[len(topic.Content)-1])
		} else {
			fmt.Println("empty")
		}

	}
}

func (b *Broker) scanForPing() {
	// now := time.Now()
	// for client, _ := range b.Clients {
	// 	// everything went right before
	// 	if client.Ping.IsPingSent && !client.Ping.IsPong {
	// 		if time.Since(client.Ping.LastPing) > b.Parameters.Ping.IntervalAfterPing {
	// 			client.SendPing()
	// 		} else {
	// 			continue
	// 		}
	// 	} else if {

	// 	}
	// }
}

// !! Ping sent
// time
// is Pong false

// !! check 1
// if time before => nothing

// !! check 2
// if time after => resend
// retry ++
// OR
// kill if retry > max_retry

// !! if get pong
// isPong true
// is PingSent true
// retry 0

// Connection to topic
// send message key/value

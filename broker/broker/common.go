package broker

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
)

var broker *Broker

type PingInfo struct {
	LastPing   time.Time `json:"time"`
	IsPingSent bool      `json:"is_ping_sent"`
	IsPong     bool      `json:"is_pong"`
	Retry      int       `json:"retry"`
}

type Watcher struct {
	Conn *websocket.Conn
	Ping PingInfo
	m    *sync.Mutex
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
	m              *sync.Mutex
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
	Watcher    map[*Watcher]bool
	Topics     map[string]*Topic
	Parameters BrokerParameters
	m          *sync.Mutex
}

func NewBroker() *Broker {
	broker = &Broker{
		Clients: map[*Client]bool{},
		Watcher: map[*Watcher]bool{},
		Topics:  map[string]*Topic{},
		Parameters: BrokerParameters{
			Ping: PingParameter{
				IntervalAfterPing: time.Second * 1,
				IntervalAfterPong: time.Second * 5,
				MaxRetry:          3,
			},
		},
		m: &sync.Mutex{},
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

func (b *Broker) LaunchLoop(ctx context.Context) {
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
	for topicName, topic := range b.Topics {
		// logrus.Infof("==> %s : consmrs : %d / queue / %d \n", topicName, len(topic.ConsumerCients), len(topic.Content))
		if len(topic.Content) > 0 {
			logrus.Warning("-> scanning : LAST CONTENT : ", topic.Content[len(topic.Content)-1])
			topic.SendJobToAvailableClient(topicName)
		} else {
			logrus.Errorf("topic %s EMPTY", topicName)
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

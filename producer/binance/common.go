package binance

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"producer/utils"
	binance_dto "shared/binance/dto"
	"shared/config"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
)

var binanceconnection *BinanceConnection

type Connection struct {
	url    url.URL
	config *websocket.Config
	conn   *websocket.Conn
}

type BinanceConnection struct {
	binance *Connection
	broker  *Connection
}

const (
	subscribeId   = 1
	unSubscribeId = 2
)

type RequestParams struct {
	Id     int      `json:"id"`
	Method string   `json:"method"`
	Params []string `json:"params"`
}

var (
	binanceUrl = url.URL{Scheme: "wss", Host: "stream.binance.com:443", Path: "/ws"}
)

func NewConn() *BinanceConnection {
	conf := config.Getenv()

	binanceConnection, err := ConnectToSocket(binanceUrl)
	if err != nil {
		log.Fatal("error with binanceConnection")
		return nil
	}

	brokerUrl := url.URL{Scheme: "ws", Host: fmt.Sprintf("%s:%s", conf.BrokerHost, conf.BrokerPort), Path: "/ws"}
	brokerConnection, err := ConnectToSocket(brokerUrl)
	if err != nil {
		log.Fatal("error with brokerConnection")
	}

	binanceconnection = &BinanceConnection{
		binance: binanceConnection,
		broker:  brokerConnection,
	}
	return binanceconnection
}

func ConnectToSocket(url url.URL) (*Connection, error) {
	config, err := websocket.NewConfig(url.String(), "http://localhost")
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	conn, err := config.DialContext(ctx)
	if err != nil {
		return nil, err
	}
	return &Connection{
		url:    url,
		config: config,
		conn:   conn,
	}, nil
}

func (c *BinanceConnection) GoListen(topic string, ctx context.Context) {

	message := RequestParams{
		Id:     subscribeId,
		Method: "SUBSCRIBE",
		Params: []string{
			// "btcusdt@aggTrade",
			"ethusdt@aggTrade",
			// "btcusdt@depth",
		},
	}
	//log.Println(message)
	b, err := json.Marshal(message)
	if err != nil {
		log.Fatal("Failed to JSON Encode trade topics")
		// return err
	}
	c.conn.Write(b)
	go func() {
		defer c.conn.Close()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				var response = make([]byte, 1024)
				n, err := c.conn.Read(response)
				logrus.Info("--> N ", n)
				if err != nil {
					panic(err)
				}
				c.handleBinanceMessage(response, n)
			}
		}
	}()
}

func (c *BinanceConnection) handleBinanceMessage(response []byte, responseLength int) error {
	fmt.Println("=> ", string(response[:responseLength]))
	logrus.Infof("%d-> %s\n", string(response[:responseLength]))

	var binanceDto binance_dto.BinanceAggTradeDto
	err := json.Unmarshal(response[:responseLength], &binanceDto)
	if err != nil {
		return err
	}
	utils.PrettyDisplay(binanceDto)
	logrus.Warn("---->", binanceDto.PriceChange)

	// shared.CustomBodyValidator()
	// mb_Conn := message_broker.GetBinanceConnection()

	// mb_Conn.Produce(topic, []byte("message from the producer"))
	// time.Sleep(time.Second)
	return nil
}

// producerName := uuid.New()
// topic := conf.BrokerTopic
// mb_conn := message_broker.NewConn(u, origin)
// dummy.GoLoopProducer(producerName.String(), topic, mb_conn.Produce, time.Second*2, ctx)

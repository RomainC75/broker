package binance

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	binance_dto "shared/binance/dto"
	message_broker "shared/broker"
	"shared/config"
	"shared/utils"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
)

var producerconnection *ProducerConnection

type BinanceConnection struct {
	url    url.URL
	config *websocket.Config
	conn   *websocket.Conn
}

type ProducerConnection struct {
	binance *BinanceConnection
	broker  *message_broker.Connection
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

func NewConn() *ProducerConnection {
	conf := config.Getenv()

	binanceConnection, err := ConnectToSocket(binanceUrl)
	if err != nil {
		log.Fatal("error with binanceConnection")
		return nil
	}

	brokerUrl := url.URL{Scheme: "ws", Host: fmt.Sprintf("%s:%s", conf.BrokerHost, conf.BrokerPort), Path: "/ws"}
	brokerConnection := message_broker.NewConn(brokerUrl, "http://localhost")
	if err != nil {
		log.Fatal("error with brokerConnection")
	}

	producerconnection = &ProducerConnection{
		binance: binanceConnection,
		broker:  brokerConnection,
	}
	return producerconnection
}

func ConnectToSocket(url url.URL) (*BinanceConnection, error) {
	config, err := websocket.NewConfig(url.String(), "http://localhost")
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	conn, err := config.DialContext(ctx)
	if err != nil {
		return nil, err
	}
	return &BinanceConnection{
		url:    url,
		config: config,
		conn:   conn,
	}, nil
}

func (c *ProducerConnection) GoListen(topic string, ctx context.Context) {

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
	}
	c.binance.conn.Write(b)
	go func() {
		defer c.binance.conn.Close()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				var response = make([]byte, 1024)
				n, err := c.binance.conn.Read(response)
				logrus.Info("--> N ", n)
				if err != nil {
					panic(err)
				}
				c.handleBinanceMessage(response, n)
			}
		}
	}()
}

func (c *ProducerConnection) handleBinanceMessage(response []byte, responseLength int) error {
	fmt.Println("=> ", string(response[:responseLength]))
	logrus.Infof("%d-> %s\n", string(response[:responseLength]))

	var binanceDto binance_dto.BinanceAggTradeDto
	err := json.Unmarshal(response[:responseLength], &binanceDto)
	if err != nil {
		return err
	}
	utils.PrettyDisplay("binance DTO", binanceDto)
	logrus.Warn("---->", binanceDto.PriceChange)

	// shared.CustomBodyValidator()
	// mb_Conn := message_broker.GetProducerConnection()

	// mb_Conn.Produce(topic, []byte("message from the producer"))
	// time.Sleep(time.Second)
	return nil
}

// producerName := uuid.New()
// topic := conf.BrokerTopic
// mb_conn := message_broker.NewConn(u, origin)
// dummy.GoLoopProducer(producerName.String(), topic, mb_conn.Produce, time.Second*2, ctx)

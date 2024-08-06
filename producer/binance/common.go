package binance

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"producer/utils"
	binance_dto "shared/binance/dto"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
)

var binanceconnection *BinanceConnection

type BinanceConnection struct {
	url    url.URL
	config *websocket.Config
	conn   *websocket.Conn
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
	u = url.URL{Scheme: "wss", Host: "stream.binance.com:443", Path: "/ws"}
)

func NewConn() *BinanceConnection {

	config, err := websocket.NewConfig(u.String(), "http://localhost")
	if err != nil {
		log.Fatal("error with config: ", err.Error())
	}
	ctx := context.Background()
	conn, err := config.DialContext(ctx)
	if err != nil {
		log.Fatal("error trying to dial: ", err.Error())
	}

	binanceconnection = &BinanceConnection{
		url:    u,
		config: config,
		conn:   conn,
	}
	return binanceconnection
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

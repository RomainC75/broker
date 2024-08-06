package binance

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	binance_dto "producer/binance/dto"
	"producer/utils"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
)

var connection *Connection

type Connection struct {
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

func NewConn() *Connection {

	config, err := websocket.NewConfig(u.String(), "http://localhost")
	if err != nil {
		log.Fatal("error with config: ", err.Error())
	}
	ctx := context.Background()
	conn, err := config.DialContext(ctx)
	if err != nil {
		log.Fatal("error trying to dial: ", err.Error())
	}

	connection = &Connection{
		url:    u,
		config: config,
		conn:   conn,
	}
	return connection
}

func (c *Connection) GoListen(topic string, ctx context.Context) {

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
		i := 0

		for {
			select {
			case <-ctx.Done():
				return
			default:
				var response = make([]byte, 1_000_000)
				n, err := c.conn.Read(response)
				logrus.Info("--> N ", n)
				if err != nil {
					panic(err)
				}
				fmt.Println("=> ", string(response[:n]))
				logrus.Infof("%d-> %s\n", i, string(response[:n]))

				var binanceDto binance_dto.BinanceAggTradeDto
				err = json.Unmarshal(response[:n], &binanceDto)
				if err != nil {
					logrus.Warn(err.Error())
					continue
				}
				utils.PrettyDisplay(binanceDto)
				logrus.Warn("---->", binanceDto.PriceChange)

				// shared.CustomBodyValidator()
				// mb_Conn := message_broker.GetConnection()

				// mb_Conn.Produce(topic, []byte("message from the producer"))
				// time.Sleep(time.Second)
				i++
			}

		}
	}()
}

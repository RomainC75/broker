package broker

import (
	"net/url"

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
	u      = url.URL{Scheme: "wss", Host: "stream.binance.com:443", Path: "/ws"}
	origin = "http://localhost"
)

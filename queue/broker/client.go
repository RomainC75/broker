package broker

import (
	"time"

	"golang.org/x/net/websocket"
)

func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		Conn: conn,
		Ping: PingInfo{
			LastPing:   time.Now(),
			IsPingSent: false,
			IsPong:     false,
		},
	}
}

package broker

import (
	"sync"
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
		IsAvailable: false,
		m:           &sync.Mutex{},
	}
}

func (c *Client) SetIsAvailable(isAvailable bool) {
	c.m.Lock()
	c.IsAvailable = isAvailable
	c.m.Unlock()
}

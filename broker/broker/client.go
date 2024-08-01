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
		IsAvailable: true,
		m:           &sync.Mutex{},
	}
}

func (c *Client) Close() {
	c.Conn.Close()
}

func (c *Client) SetIsAvailable(isAvailable bool) {
	c.m.Lock()
	c.IsAvailable = isAvailable
	c.m.Unlock()
}

func (c *Client) SendPing() {
	now := time.Now()
	c.m.Lock()
	c.Ping.LastPing = now
	c.Ping.IsPong = false
	c.Ping.IsPingSent = true
	c.m.Unlock()
}

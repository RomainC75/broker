package broker

import (
	"time"

	"golang.org/x/net/websocket"
)

func NewWatcher(conn *websocket.Conn) *Watcher {
	return &Watcher{
		Conn: conn,
	}
}

func LaunchWatcher() {

	for {

		time.Sleep(time.Second * 2)
	}
}

func (b *Broker) AddWatcher(conn *websocket.Conn) {
	watcher := NewWatcher(conn)
	b.m.Lock()
	b.Watcher[watcher] = true
	b.m.Unlock()
}

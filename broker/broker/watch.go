package broker

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
)

func NewWatcher(conn *websocket.Conn) *Watcher {
	return &Watcher{
		Conn: conn,
		m:    &sync.Mutex{},
	}
}

func (b *Broker) LaunchWatcherLoop(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				b.BroadcastInfosToWatchers()
				time.Sleep(time.Second * 2)
			}
		}
	}()
}

func (b *Broker) AddWatcher(conn *websocket.Conn) {
	var wg sync.WaitGroup
	wg.Add(1)
	watcher := NewWatcher(conn)
	b.m.Lock()
	b.Watcher[watcher] = true
	b.m.Unlock()
	// TODO : add a listener ??
	wg.Wait()
}

func (b *Broker) BroadcastInfosToWatchers() {
	dataToSent := ToTopicsDtoToSend(b)
	// fmt.Println("watcher looop")
	// utils.PrettyDisplay("PRETTY : ", dataToSent)
	by, err := json.Marshal(dataToSent)
	if err != nil {
		logrus.Error(err.Error())
	}

	for w := range b.Watcher {
		// w.Conn.Write(by)
		_, err := w.Conn.Write(by)
		if err != nil {
			fmt.Println("error trying to write socket ", err.Error())
		}
	}
}

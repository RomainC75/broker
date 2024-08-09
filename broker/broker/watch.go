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
				time.Sleep(time.Millisecond * time.Duration(b.Parameters.Watcher.BrokerWatcherFrequenceMs))
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

func (b *Broker) RemoveWatcher(watchers []*Watcher) {
	if len(watchers) == 0 {
		return
	}
	b.m.Lock()
	for _, w := range watchers {
		b.Watcher[w] = false
	}
	b.m.Unlock()
}

func (b *Broker) BroadcastInfosToWatchers() {
	dataToSent := ToTopicsDtoToSend(b, b.Parameters.Watcher)
	by, err := json.Marshal(dataToSent)
	if err != nil {
		logrus.Error(err.Error())
	}

	watchersToRemove := []*Watcher{}

	for w := range b.Watcher {
		_, err := w.Conn.Write(by)
		if err != nil {
			fmt.Println("error trying to write socket ", err.Error())
			watchersToRemove = append(watchersToRemove, w)
		}
	}
	b.RemoveWatcher(watchersToRemove)
}

package dummy

import (
	"context"
	"fmt"
	"time"
)

func GoLoopProducer(producerName string, topicName string, fn func(string, []byte), t time.Duration, ctx context.Context) {
	go func() {
		index := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				fn(topicName, []byte(fmt.Sprintf("%s : message - %d", producerName, index)))
				index++
				time.Sleep(t)
			}
		}
	}()
}

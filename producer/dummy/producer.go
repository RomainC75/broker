package dummy

import (
	"context"
	"fmt"
	"time"
)

func GoLoopProducer(topicName string, fn func(string, []byte), t time.Duration, ctx context.Context) {
	go func() {
		index := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				fn(topicName, []byte(fmt.Sprintf("message - %d", index)))
				index++
				time.Sleep(t)
			}
		}
	}()
}

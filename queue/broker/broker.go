package broker

import "time"

func (b *Broker) GoLauchBroker() {
	go func() {
		for {
			for c, _ := range b.Clients {
				c.Conn.Write([]byte("broadcast"))
			}
			time.Sleep(time.Second * 3)
		}
	}()
}

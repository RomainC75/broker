package broker

import (
	service "broker/services"
	"encoding/json"
	"errors"
	"shared/broker_dto"
	"sync"

	"github.com/sirupsen/logrus"
)

const QUEUE_LENGTH = 2000

func NewTopic() *Topic {
	return &Topic{
		Queue:          service.NewQueue[Message](QUEUE_LENGTH),
		ConsumerCients: map[*Client]bool{},
		m:              &sync.Mutex{},
	}
}

func (b *Broker) addClientToTopic(topic string, client *Client) error {
	if entry, ok := b.Topics[topic]; ok {
		entry.ConsumerCients[client] = true
		b.Topics[topic] = entry
		return nil
	} else {
		return errors.New("not topic found")
	}
}

func (b *Broker) removeClientFromTopic(topic string, client *Client) error {
	if entry, ok := b.Topics[topic]; ok {
		// many clients ?
		if len(entry.ConsumerCients) > 1 {
			// yes : remove one
			delete(entry.ConsumerCients, client)
			// no : check messages
		} else if entry.Queue.GetSize() > 1 {
			// messages ?  =>  remove client
			delete(entry.ConsumerCients, client)
		} else {
			// no message => remove topic
			delete(b.Topics, topic)
		}
		return nil
	} else {
		return errors.New("not topic found")
	}
}

func (t *Topic) SendJobToAvailableClient(topicName string) {
	for c := range t.ConsumerCients {
		if c.IsAvailable {
			// var nextJob []byte
			nextJob, id, _ := t.Queue.GetFirstValueAndToHandling()

			contentB, err := json.Marshal(nextJob)
			if err != nil {
				logrus.Warn(err.Error())
			}
			// nextJob = jobContent.Value
			message := broker_dto.Message{
				Topic:      topicName,
				ActionCode: broker_dto.SendJob,
				// Offset:     i,
				Id:      id,
				Content: contentB,
			}
			messageB, err := json.Marshal(message)
			if err != nil {
				logrus.Errorf("Error trying to marshall message in topic %s\n", topicName)
			}
			_, err = c.Conn.Write(messageB)
			if err != nil {
				logrus.Errorf("Error trying to send message in topic %s\n", topicName)
			}
			logrus.Infof("message SENT in topic %s\n", topicName)
			break
		}
	}
}

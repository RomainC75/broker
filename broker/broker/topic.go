package broker

import (
	"encoding/json"
	"errors"
	"shared/broker_dto"
	"sync"

	"github.com/sirupsen/logrus"
)

func NewTopic() *Topic {
	return &Topic{
		Content:        []Message{},
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
			entry.ConsumerCients[client] = false
			// no : check messages
		} else if len(entry.Content) > 1 {
			// messages ?  =>  remove client
			entry.ConsumerCients[client] = false
		} else {
			// no message => remove topic
			delete(b.Topics, topic)
		}
		return nil
	} else {
		return errors.New("not topic found")
	}
}

// ====TOPIC====
func (t *Topic) SendJobToAvailableClient(topicName string) {
	// logrus.Warn("TRYING TO SEND TO ", topicName, len(t.ConsumerCients))
	for c := range t.ConsumerCients {
		// logrus.Infof("is client available : %t\n", c.IsAvailable)
		if c.IsAvailable {
			var nextJob []byte
			for i, jobContent := range t.Content {
				// logrus.Warn("potential job : ")
				// utils.PrettyDisplay("job", jobContent)
				if !jobContent.IsSent && !jobContent.IsHandled {
					// logrus.Info("FOUND 1 CLIENT AVAILABLE")
					nextJob = jobContent.Value
					t.Content[i].IsSent = true
					message := broker_dto.Message{
						Topic:      topicName,
						ActionCode: broker_dto.SendJob,
						Offset:     i,
						Content:    nextJob,
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
	}
}

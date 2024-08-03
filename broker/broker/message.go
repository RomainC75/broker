package broker

import (
	"errors"
	"fmt"
	"shared/broker_dto"

	"github.com/sirupsen/logrus"
)

func (b *Broker) addMessage(message broker_dto.Message) {
	isExist := b.isTopicExists(message.Topic)
	if !isExist {
		fmt.Println("does not exist ")
		b.Topics[message.Topic] = NewTopic()
	}
	fmt.Println("LEN : ", len(b.Topics[message.Topic].Content))
	newMessage := Message{
		Value: message.Content,
	}
	topic := b.Topics[message.Topic]
	topic.Content = append(topic.Content, newMessage)
	b.Topics[message.Topic] = topic

	topic = b.Topics[message.Topic]
	logrus.Warn("CHECK : ", len(topic.Content))

}

func (b *Broker) SetJobToAccepted(topic string, offset int) error {
	data, ok := b.Topics[topic]
	if !ok {
		return errors.New("topic not found ")
	}
	data.m.Lock()
	data.Content[offset].IsHandled = true
	data.m.Unlock()
	return nil
}

func (b *Broker) isTopicExists(name string) bool {
	_, ok := b.Topics[name]
	return ok
}

func (b *Broker) AdaptReaderIndex(topic string) error {
	data, ok := b.Topics[topic]
	if !ok {
		return errors.New("topic not found ")
	}
	for i := data.ReaderIndex; i < len(data.Content); i++ {
		if !data.Content[i].IsHandled {
			data.m.Lock()
			data.ReaderIndex = i
			data.m.Unlock()
			return nil
		}
	}
	return nil
}

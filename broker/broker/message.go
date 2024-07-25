package broker

import (
	"fmt"
	"shared/broker_dto"
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
}

func (b *Broker) isTopicExists(name string) bool {
	_, ok := b.Topics[name]
	return ok
}

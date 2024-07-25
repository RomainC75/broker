package broker

import "errors"

func NewTopic() Topic {
	return Topic{
		Content:        []Message{},
		ConsumerCients: map[*Client]bool{},
	}
}

func (b *Broker) addClientToTopic(topic string, client *Client) error {
	if entry, ok := b.Topics[topic]; ok {
		entry.ConsumerCients[client] = true
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

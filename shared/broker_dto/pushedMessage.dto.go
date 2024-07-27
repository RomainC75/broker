package broker_dto

import "encoding/json"

type ActionCode int

const (
	UnSubscribe ActionCode = 0
	Subscribe   ActionCode = 1
	// to Broker
	SendMessage ActionCode = 2
	Ping        ActionCode = 3
	Pong        ActionCode = 4
	IsAvailable ActionCode = 5
	// send Job to Consumer
	SendJob ActionCode = 6
)

type Message struct {
	Topic      string     `json:"topic"`
	ActionCode ActionCode `json:"request"`
	Content    []byte     `json:"content"`
}

type IsAvailableContent struct {
	IsAvailable bool `json:"is_available"`
}

func GetIsAvailableMessage(isAvailable bool) (Message, error) {
	isAvailableContent := IsAvailableContent{
		IsAvailable: isAvailable,
	}
	isAvailableB, err := json.Marshal(isAvailableContent)
	if err != nil {
		return Message{}, err
	}
	return Message{
		ActionCode: IsAvailable,
		Content:    isAvailableB,
	}, nil
}

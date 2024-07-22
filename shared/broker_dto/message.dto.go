package broker_dto

type SubscribeMessage int

const (
	UnSubscribe SubscribeMessage = 0
	Subscribe   SubscribeMessage = 1
)

type TopicSelection struct {
	Topic   string `json:"topic"`
	Request SubscribeMessage
}

type IsAvailable struct {
	Type string
}

type Message struct {
	Type    string   `json:"type"`
	Content struct{} `json:"content"`
}

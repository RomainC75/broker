package broker_dto

type ActionCode int

const (
	UnSubscribe ActionCode = 0
	Subscribe   ActionCode = 1
	SendMessage ActionCode = 2
	Ping        ActionCode = 3
	Pong        ActionCode = 4
	IsAvailable ActionCode = 5
)

type Message struct {
	Topic      string     `json:"topic"`
	ActionCode ActionCode `json:"request"`
	Content    []byte     `json:"content"`
}

type IsAvailableContent struct {
	IsAvailable bool `json:"is_available"`
}

package broker_dto

type ActionCode int

const (
	UnSubscribe ActionCode = 0
	Subscribe   ActionCode = 1
	SendMessage ActionCode = 2
)

type Message struct {
	Topic      string     `json:"topic"`
	ActionCode ActionCode `json:"request"`
	Content    struct{}   `json:"content"`
}

type IsAvailable struct {
	Type string
}

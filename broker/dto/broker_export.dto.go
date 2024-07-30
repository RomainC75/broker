package broker_client_dto

import (
	"broker/broker"
)

// type PingInfo struct {
// 	LastPing   time.Time `json:"time"`
// 	IsPingSent bool      `json:"is_ping_sent"`
// 	IsPong     bool      `json:"is_pong"`
// 	Retry      int       `json:"retry"`
// }

type MessageDto struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	IsSent    bool   `json:"is_sent"`
	IsHandled bool   `json:"is_handled"`
}

type ClientDto struct {
	Ping        broker.PingInfo `json:"ping"`
	Topics      []string        `json:"topics"`
	IsAvailable bool            `json:"is_available"`
}

type TopicDto struct {
	Content        []MessageDto `json:"content"`
	ConsumerCients []ClientDto  `json:"consumer_clients"`
	ReaderIndex    int          `json:"reader_index"`
}

func ToTopicsDtoToSend(broker *broker.Broker) map[string]TopicDto {
	topics := make(map[string]TopicDto)

	return topics
}

func ToPingInfo(pingInfo broker.PingInfo) broker.PingInfo {
	return broker.PingInfo{
		LastPing:   pingInfo.LastPing,
		IsPingSent: pingInfo.IsPingSent,
		IsPong:     pingInfo.IsPong,
		Retry:      pingInfo.Retry,
	}
}

func ToConsumerClients(rawClients map[*broker.Client]bool) []ClientDto {
	clients := []ClientDto{}
	for rawClient := range rawClients {
		clients = append(clients, ClientDto{
			Ping:        rawClient.Ping,
			Topics:      rawClient.Topics,
			IsAvailable: rawClient.IsAvailable,
		})
	}
	return clients
}

func ToTopicDto(topic broker.Topic) TopicDto {
	return TopicDto{
		Content:        ToMessageDto(topic.Content),
		ConsumerCients: ToConsumerClients(topic.ConsumerCients),
		ReaderIndex:    topic.ReaderIndex,
	}
}

func ToMessageDto(messages []broker.Message) []MessageDto {
	dtoMessages := []MessageDto{}
	for _, message := range messages {
		dtoMessages = append(dtoMessages, MessageDto{
			Key:       string(message.Key),
			Value:     string(message.Value),
			IsSent:    message.IsSent,
			IsHandled: message.IsHandled,
		})
	}
	return dtoMessages
}

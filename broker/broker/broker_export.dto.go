package broker

// type PingInfo struct {
// 	LastPing   time.Time `json:"time"`
// 	IsPingSent bool      `json:"is_ping_sent"`
// 	IsPong     bool      `json:"is_pong"`
// 	Retry      int       `json:"retry"`
// }

type MessageDto struct {
	Index     int    `json:"index"`
	Key       string `json:"key"`
	Value     string `json:"value"`
	IsSent    bool   `json:"is_sent"`
	IsHandled bool   `json:"is_handled"`
}

type ClientDto struct {
	Ping        PingInfo `json:"ping"`
	Topics      []string `json:"topics"`
	IsAvailable bool     `json:"is_available"`
}

type TopicDto struct {
	Content        []MessageDto `json:"content"`
	ConsumerCients []ClientDto  `json:"consumer_clients"`
	ReaderIndex    int          `json:"reader_index"`
}

type TopicMapDto map[string]TopicDto

func ToTopicsDtoToSend(broker *Broker) map[string]TopicDto {
	topics := make(map[string]TopicDto)
	for topicName, topic := range broker.Topics {
		topics[topicName] = ToTopicDto(*topic)
	}
	return topics
}

func ToPingInfo(pingInfo PingInfo) PingInfo {
	return PingInfo{
		LastPing:   pingInfo.LastPing,
		IsPingSent: pingInfo.IsPingSent,
		IsPong:     pingInfo.IsPong,
		Retry:      pingInfo.Retry,
	}
}

func ToConsumerClients(rawClients map[*Client]bool) []ClientDto {
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

func ToTopicDto(topic Topic) TopicDto {
	return TopicDto{
		Content:        ToMessageDto(topic.Content),
		ConsumerCients: ToConsumerClients(topic.ConsumerCients),
		ReaderIndex:    topic.ReaderIndex,
	}
}

func ToMessageDto(messages []Message) []MessageDto {
	dtoMessages := []MessageDto{}
	for index, message := range messages {
		dtoMessages = append(dtoMessages, MessageDto{
			Index:     index,
			Key:       string(message.Key),
			Value:     string(message.Value),
			IsSent:    message.IsSent,
			IsHandled: message.IsHandled,
		})
	}
	return dtoMessages
}

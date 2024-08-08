package config

type Config struct {
	BrokerPort                       string
	BrokerHost                       string
	BrokerTopic                      string
	BrokerWatcherFrequenceMs         int
	Tickers                          string
	BrockerWatcherTopicContentLength int
	Db                               DB
}

type DB struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

package config

type Config struct {
	BrokerPort               string
	BrokerHost               string
	BrokerTopic              string
	BrokerWatcherFrequenceMs int
	Tickers                  string
}

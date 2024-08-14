package config

type Config struct {
	BrokerPort                       string
	BrokerHost                       string
	BrokerTopic                      string
	BrokerWatcherFrequenceMs         int
	Tickers                          []string
	BrockerWatcherTopicContentLength int
	Db                               DB
	Azure                            Azure
	Redis                            Redis
}

type DB struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

type Azure struct {
	AudienceId string
	TenantId   string
}

type Redis struct {
	Port     int
	Host     string
	RefreshM int
}

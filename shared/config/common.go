package config

import (
	"log"
	"os"
	config_utils "shared/config/utils"
	"strconv"

	"github.com/joho/godotenv"
)

var config *Config

func Getenv() *Config {
	return config
}

func SetEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("not .env found ! ")
	}
	config = &Config{}

	(*config).BrokerHost = os.Getenv("BROKER_HOST")
	(*config).BrokerPort = os.Getenv("BROKER_PORT")
	(*config).BrokerTopic = os.Getenv("BROKER_TOPIC")
	(*config).Tickers = config_utils.SeparateTickers(os.Getenv("TICKERS"))

	interval := os.Getenv("BROKER_WATCHER_INTERVAL_MS")
	intervalInt, err := strconv.Atoi(interval)
	if err != nil {
		log.Fatal("BROKER_WATCHER_INTERVAL_MS not valid in .env file !!")
	}
	(*config).BrokerWatcherFrequenceMs = intervalInt

	topicContentLength := os.Getenv("BROKER_WATCHER_TOPIC_CONTENT_LENGTH")
	brockerWatcherTopicContentLength, err := strconv.Atoi(topicContentLength)
	if err != nil {
		log.Fatal("BROKER_WATCHER_TOPIC_CONTENT_LENGTH not valid in .env file !!")
	}
	(*config).BrockerWatcherTopicContentLength = brockerWatcherTopicContentLength

	(*config).Db.Host = os.Getenv("POSTGRES_HOST")
	pgPort := os.Getenv("POSTGRES_PORT")
	pgPortInt, err := strconv.Atoi(pgPort)
	if err != nil {
		log.Fatal("POSTGRES_PORT not valid in .env file !!")
	}
	(*config).Db.Port = pgPortInt
	(*config).Db.User = os.Getenv("POSTGRES_USER")
	(*config).Db.Password = os.Getenv("POSTGRES_PASSWORD")
	(*config).Db.DbName = os.Getenv("POSTGRES_DB_NAME")

	(*config).Azure.AudienceId = os.Getenv("VITE_API_SSO_AUDIENCE_ID")
	(*config).Azure.TenantId = os.Getenv("VITE_API_SSO_TENANT_ID")
}

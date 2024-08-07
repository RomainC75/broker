package config

import (
	"log"
	"os"
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
	(*config).Tickers = os.Getenv("TICKERS")

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
}

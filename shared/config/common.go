package config

import (
	"log"
	"os"

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
}

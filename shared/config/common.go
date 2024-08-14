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

	// BROKER
	(*config).BrokerHost = os.Getenv("BROKER_HOST")
	(*config).BrokerPort = os.Getenv("BROKER_PORT")
	(*config).BrokerTopic = os.Getenv("BROKER_TOPIC")
	(*config).Tickers = config_utils.SeparateTickers(os.Getenv("TICKERS"))
	(*config).BrokerWatcherFrequenceMs = getIntEnvVariable("BROKER_WATCHER_INTERVAL_MS")
	(*config).BrockerWatcherTopicContentLength = getIntEnvVariable("BROKER_WATCHER_TOPIC_CONTENT_LENGTH")

	// POSTGRES
	(*config).Db.Host = os.Getenv("POSTGRES_HOST")
	(*config).Db.Port = getIntEnvVariable("POSTGRES_PORT")
	(*config).Db.User = os.Getenv("POSTGRES_USER")
	(*config).Db.Password = os.Getenv("POSTGRES_PASSWORD")
	(*config).Db.DbName = os.Getenv("POSTGRES_DB_NAME")

	// AZURE
	(*config).Azure.AudienceId = os.Getenv("VITE_API_SSO_AUDIENCE_ID")
	(*config).Azure.TenantId = os.Getenv("VITE_API_SSO_TENANT_ID")

	// REDIS
	(*config).Redis.Port = getIntEnvVariable("REDIS_PORT")
	(*config).Redis.RefreshM = getIntEnvVariable("REDIS_REFRESH_M")
	(*config).Redis.Host = os.Getenv("REDIS_HOST")
}

func getIntEnvVariable(varName string) int {
	strVar := os.Getenv(varName)
	strVarInt, err := strconv.Atoi(strVar)
	if err != nil {
		log.Fatalf("%s is not a valid for %s in the .env file !", strVar, varName)
	}
	return strVarInt
}

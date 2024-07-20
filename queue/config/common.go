package config

import (
	"fmt"
	"log"
	"os"
	"queue/utils"
	"reflect"

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

	val := reflect.ValueOf(*config)
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		envFormat := utils.ToEnvString(field.Name)
		value := os.Getenv(envFormat)

		if len(value) == 0 {
			log.Fatalf("%s not found in .env", field.Name)
		}
		fmt.Printf("name : %s value : %s\n", field.Name, value)
	}
}

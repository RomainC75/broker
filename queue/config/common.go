package config

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

var config *Config

func SetEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("not .env found ! ")
	}
	config = &Config{}

	val := reflect.ValueOf(*config)
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)

		value := os.Getenv(field.Name)
		if len(value) == 0 {
			log.Fatalf("%s not found in .env", field.Name)
		}
		fmt.Printf("name : %s value : %s\n", field.Name, value)
	}

}

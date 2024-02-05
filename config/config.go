package config

import (
	"os"
)

type Config struct {
	Host        string
	User        string
	Password    string
	ConsumerKey string
}

func GetConfig() Config {
	return Config{
		Host:        os.Getenv("OBP_HOSTNAME"),
		User:        os.Getenv("OBP_USERNAME"),
		Password:    os.Getenv("OBP_PASSWORD"),
		ConsumerKey: os.Getenv("OBP_CONSUMER_KEY"),
	}
}

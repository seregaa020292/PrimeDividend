package config

import (
	"os"
)

type Config struct {
	Env     string
	Tinkoff struct {
		AuthToken string
	}
	Redis struct {
		Host string
		Port string
	}
}

func NewConfig() Config {
	config := Config{}

	config.Env = os.Getenv("APP_ENV")
	config.Tinkoff.AuthToken = os.Getenv("TINKOFF_AUTH_TOKEN")

	config.Redis.Host = os.Getenv("REDIS_HOST")
	config.Redis.Port = os.Getenv("REDIS_PORT")

	return config
}

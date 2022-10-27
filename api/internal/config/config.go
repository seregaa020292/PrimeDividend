package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App      App
	Mailer   Mailer
	Postgres Postgres
	Redis    Redis
	Networks Networks
	Jwt      Jwt
	Tinkoff  Tinkoff
}

var (
	instance Config
	once     sync.Once
)

func GetConfig() Config {
	once.Do(func() {
		log.Println("Start Config")

		instance = Config{}

		if err := cleanenv.ReadEnv(&instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Println(help)
			log.Fatalln(err)
		}
	})
	return instance
}

package config

import (
	"io/ioutil"
	"log"
	"strings"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type App struct {
	Env        string `env:"APP_ENV" env-default:"development"`
	SiteOrigin string `env:"SITE_ORIGIN" env-default:"http://localhost"`
	LogLevel   string `env:"LOG_LEVEL" env-default:"debug"`
}

func (a App) IsDevelopment() bool {
	return a.Env == "development"
}

type Mailer struct {
	Host         string `env:"MAILER_HOST"`
	Port         int    `env:"MAILER_PORT"`
	Username     string `env:"MAILER_USERNAME"`
	PasswordFile string `env:"MAILER_PASSWORD_FILE"`
	Encryption   string `env:"MAILER_ENCRYPTION"`
	FromEmail    string `env:"MAILER_FROM_EMAIL"`
}

func (m Mailer) Password() string {
	file, _ := ioutil.ReadFile(m.PasswordFile)
	return strings.TrimSpace(string(file))
}

type Postgres struct {
	Username     string `env:"DB_USER" env-required:"true"`
	PasswordFile string `env:"DB_PASSWORD_FILE" env-required:"true"`
	Host         string `env:"DB_HOST" env-required:"true"`
	Port         int    `env:"DB_PORT" env-required:"true"`
	Database     string `env:"DB_NAME" env-required:"true"`
}

func (p Postgres) Password() string {
	file, _ := ioutil.ReadFile(p.PasswordFile)
	return strings.TrimSpace(string(file))
}

type Config struct {
	App      App
	Mailer   Mailer
	Postgres Postgres
}

var (
	instance Config
	once     sync.Once
)

func GetConfig() Config {
	once.Do(func() {
		log.Println("Start config")

		instance = Config{}

		if err := cleanenv.ReadEnv(&instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Println(help)
			log.Fatalln(err)
		}
	})
	return instance
}

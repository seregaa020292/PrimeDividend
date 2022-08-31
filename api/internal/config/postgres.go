package config

import (
	"io/ioutil"
	"strings"
)

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

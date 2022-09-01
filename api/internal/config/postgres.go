package config

import (
	"fmt"
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

func (p Postgres) Dsn() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s sslmode=disable password=%s",
		p.Host,
		p.Port,
		p.Username,
		p.Database,
		p.Password(),
	)
}

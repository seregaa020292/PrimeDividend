package config

import (
	"io/ioutil"
	"strings"
)

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

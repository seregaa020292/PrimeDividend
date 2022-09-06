package config

import (
	"io/ioutil"
	"net/mail"
	"strings"
)

type Mailer struct {
	Host         string `env:"MAILER_HOST"`
	Port         int    `env:"MAILER_PORT"`
	Username     string `env:"MAILER_USERNAME"`
	PasswordFile string `env:"MAILER_PASSWORD_FILE"`
	FromEmail    string `env:"MAILER_FROM_EMAIL"`
	FromName     string `env:"MAILER_FROM_NAME"`
}

func (m Mailer) From() mail.Address {
	return mail.Address{
		Name:    m.FromName,
		Address: m.FromEmail,
	}
}

func (m Mailer) Password() string {
	file, _ := ioutil.ReadFile(m.PasswordFile)
	return strings.TrimSpace(string(file))
}

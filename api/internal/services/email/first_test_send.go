package email

import (
	"fmt"
	"primedivident/internal/config"
	"primedivident/pkg/mailer"
)

type FirstTestSend struct {
	config config.Config
	mailer mailer.Sender
}

func NewFirstTestSend(config config.Config, mailer mailer.Sender) FirstTestSend {
	return FirstTestSend{
		config: config,
		mailer: mailer,
	}
}

func (s FirstTestSend) Send(data []byte) error {
	msg := mailer.NewMessage("Theme First Test", fmt.Sprintf("%s %s", "Hello Mail Sender!!!", data))
	msg.To = []string{"serega020292@mail.ru"}

	return s.mailer.Send(msg)
}

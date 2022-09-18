package email

import (
	"primedivident/internal/decorator"
	"primedivident/pkg/mailer"
)

type (
	ConfirmData struct {
		Email string
	}
	ConfirmUser decorator.Sender[ConfirmData]
)

type confirmUser struct {
	mailer mailer.Sender
}

func NewConfirmUser(mailer mailer.Sender) ConfirmUser {
	return confirmUser{
		mailer: mailer,
	}
}

func (s confirmUser) Send(data ConfirmData) error {
	msg := mailer.NewMessage(
		"Поздравляем",
		"Поздравляем Вы успешно зарегистрировались",
	)
	msg.To = []string{data.Email}

	return s.mailer.Send(msg)
}

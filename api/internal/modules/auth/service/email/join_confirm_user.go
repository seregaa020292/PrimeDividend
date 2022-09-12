package email

import (
	"fmt"

	"primedivident/pkg/mailer"
)

type JoinConfirmUser struct {
	mailer mailer.Sender
}

func NewJoinConfirmUser(mailer mailer.Sender) JoinConfirmUser {
	return JoinConfirmUser{
		mailer: mailer,
	}
}

func (s JoinConfirmUser) Send(email, token string) error {
	msg := mailer.NewMessage(
		"Подтвердите ваш адрес электронной почты",
		fmt.Sprintf("Token: %s", token),
	)
	msg.To = []string{email}

	return s.mailer.Send(msg)
}

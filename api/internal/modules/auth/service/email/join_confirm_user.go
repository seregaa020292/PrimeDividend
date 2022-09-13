package email

import (
	"fmt"

	"primedivident/internal/decorator"
	"primedivident/pkg/mailer"
)

type (
	ConfirmData struct {
		Email string
		Token string
	}
	JoinConfirmUser decorator.Sender[ConfirmData]
)

type joinConfirmUser struct {
	mailer mailer.Sender
}

func NewJoinConfirmUser(mailer mailer.Sender) JoinConfirmUser {
	return joinConfirmUser{
		mailer: mailer,
	}
}

func (s joinConfirmUser) Send(data ConfirmData) error {
	msg := mailer.NewMessage(
		"Подтвердите ваш адрес электронной почты",
		fmt.Sprintf(`Token: <a href="/">%s</a>`, data.Token),
		mailer.TextHtml,
	)
	msg.To = []string{data.Email}

	return s.mailer.Send(msg)
}

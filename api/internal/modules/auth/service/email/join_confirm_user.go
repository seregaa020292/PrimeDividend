package email

import (
	"primedivident/internal/config/consts"
	"primedivident/internal/decorator"
	"primedivident/pkg/mailer"
	"primedivident/pkg/tpl"
)

type (
	JoinData struct {
		Email string
		Token string
	}
	JoinConfirmUser decorator.Sender[JoinData]
)

type joinConfirmUser struct {
	mailer   mailer.Sender
	template tpl.Templater
}

func NewJoinConfirmUser(mailer mailer.Sender, template tpl.Templater) JoinConfirmUser {
	return joinConfirmUser{
		mailer:   mailer,
		template: template,
	}
}

func (s joinConfirmUser) Send(data JoinData) error {
	html, err := s.template.Render(consts.TemplateMailConfirmed, map[string]any{
		"token": data.Token,
	})
	if err != nil {
		return err
	}

	msg := mailer.NewMessage(
		"Подтвердите вашу почту",
		string(html),
		mailer.TextHtml,
	)
	msg.To = []string{data.Email}

	return s.mailer.Send(msg)
}

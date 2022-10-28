package email

import (
	"primedividend/api/internal/config/consts"
	"primedividend/api/internal/decorators"
	"primedividend/api/pkg/mailer"
	"primedividend/api/pkg/tpl"
)

type (
	ConfirmData struct {
		Email string
	}
	ConfirmUser decorators.Sender[ConfirmData]
)

type confirmUser struct {
	mailer   mailer.Sender
	template tpl.Templater
}

func NewConfirmUser(mailer mailer.Sender, template tpl.Templater) ConfirmUser {
	return confirmUser{
		mailer:   mailer,
		template: template,
	}
}

func (s confirmUser) Send(data ConfirmData) error {
	html, err := s.template.RenderInline(consts.TemplateMailConfirmed, nil)
	if err != nil {
		return err
	}

	msg := mailer.NewMessage("Поздравляем", html, mailer.TextHtml)
	msg.To = []string{data.Email}

	return s.mailer.Send(msg)
}

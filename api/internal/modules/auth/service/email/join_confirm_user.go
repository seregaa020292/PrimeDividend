package email

import (
	"primedividend/api/internal/config/consts"
	"primedividend/api/internal/decorators"
	"primedividend/api/pkg/mailer"
	"primedividend/api/pkg/tpl"
)

type (
	JoinData struct {
		Email string
		Token string
	}
	JoinConfirmUser decorators.Sender[JoinData]
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
	html, err := s.template.RenderInline(consts.TemplateMailToken, map[string]any{
		"token": data.Token,
	})
	if err != nil {
		return err
	}

	msg := mailer.NewMessage("Подтвердите вашу почту", html, mailer.TextHtml)
	msg.To = []string{data.Email}

	return s.mailer.Send(msg)
}

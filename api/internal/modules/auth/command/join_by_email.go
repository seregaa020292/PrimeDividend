package command

import (
	"primedivident/internal/config/consts"
	"primedivident/internal/decorator"
	"primedivident/internal/modules/auth/entity"
	"primedivident/internal/modules/auth/repository"
	"primedivident/internal/modules/auth/service/email"
	"primedivident/pkg/errorn"
	"primedivident/pkg/utils/hash"
)

type (
	Credential struct {
		Email    string
		Password string
	}
	JoinByEmail decorator.CommandHandler[Credential]
)

type joinByEmail struct {
	repository repository.Repository
	email      email.JoinConfirmUser
}

func NewJoinByEmail(
	repository repository.Repository,
	email email.JoinConfirmUser,
) JoinByEmail {
	return joinByEmail{
		repository: repository,
		email:      email,
	}
}

func (c joinByEmail) Exec(cmd Credential) error {
	hasEmail, err := c.repository.HasByEmail(cmd.Email)
	if err != nil {
		return errorn.ErrorSelect.Wrap(err)
	}
	if hasEmail {
		return errorn.ErrorExistEmail
	}

	user := entity.User{
		Email:    cmd.Email,
		Password: hash.Password(cmd.Password),
		Token:    entity.NewToken(consts.TokenTTL),
	}

	if err := c.repository.Add(user); err != nil {
		return errorn.ErrorInsert.Wrap(err)
	}

	if err := c.email.Send(email.ConfirmData{
		Email: user.Email,
		Token: user.Token.String(),
	}); err != nil {
		return errorn.ErrorSendEmail.Wrap(err)
	}

	return nil
}

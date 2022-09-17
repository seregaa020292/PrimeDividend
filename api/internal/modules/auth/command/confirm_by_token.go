package command

import (
	"primedivident/internal/decorator"
	"primedivident/internal/modules/auth/repository"
	"primedivident/internal/modules/auth/service/email"
	"primedivident/pkg/errorn"
)

type (
	ConfirmByToken decorator.CommandHandler[string]
)

type confirmByToken struct {
	repository repository.Repository
	email      email.JoinConfirmUser
}

func NewConfirmByToken(
	repository repository.Repository,
	email email.JoinConfirmUser,
) ConfirmByToken {
	return confirmByToken{
		repository: repository,
		email:      email,
	}
}

func (c confirmByToken) Exec(cmd string) error {
	if err := c.repository.Confirm(cmd); err != nil {
		return errorn.ErrorUpdate.Wrap(err)
	}

	return nil
}

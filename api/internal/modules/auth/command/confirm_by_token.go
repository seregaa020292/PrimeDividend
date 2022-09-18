package command

import (
	"github.com/google/uuid"

	"primedivident/internal/decorator"
	"primedivident/internal/modules/auth/repository"
	"primedivident/internal/modules/auth/service/email"
	"primedivident/pkg/errorn"
)

type (
	ConfirmByToken decorator.CommandHandler[uuid.UUID]
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

func (c confirmByToken) Exec(tokenValue uuid.UUID) error {
	token, err := c.repository.FindByTokenJoin(tokenValue)
	if err != nil {
		return errorn.ErrorSelect.Wrap(err)
	}

	if token.IsExpiredByNow() {
		return errorn.ErrorAccess
	}

	if err := c.repository.Confirm(token.Value); err != nil {
		return errorn.ErrorUpdate.Wrap(err)
	}

	return nil
}

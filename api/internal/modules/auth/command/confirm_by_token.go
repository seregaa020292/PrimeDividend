package command

import (
	"fmt"

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
	email      email.ConfirmUser
}

func NewConfirmByToken(
	repository repository.Repository,
	email email.ConfirmUser,
) ConfirmByToken {
	return confirmByToken{
		repository: repository,
		email:      email,
	}
}

func (c confirmByToken) Exec(tokenValue uuid.UUID) error {
	user, err := c.repository.FindByTokenJoin(tokenValue)
	if err != nil {
		return errorn.ErrSelect.Wrap(err)
	}

	if user.Token.IsExpiredByNow() {
		return errorn.ErrForbidden.Wrap(fmt.Errorf("%s", "token expired"))
	}

	if err := c.repository.Confirm(user.Token.Value); err != nil {
		return errorn.ErrUpdate.Wrap(err)
	}

	if err := c.email.Send(email.ConfirmData{Email: user.Email}); err != nil {
		return errorn.ErrSendEmail.Wrap(err)
	}

	return nil
}

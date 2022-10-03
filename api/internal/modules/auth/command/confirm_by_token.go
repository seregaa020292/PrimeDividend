package command

import (
	"github.com/google/uuid"

	"primedivident/internal/decorator"
	"primedivident/internal/modules/auth/repository"
	"primedivident/internal/modules/auth/service/email"
	"primedivident/pkg/errs"
	"primedivident/pkg/errs/errmsg"
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
		return errs.NotFound.Wrap(err, errmsg.CouldNotBeFound)
	}

	if err := user.Token.ErrorIsExpiredByNow(); err != nil {
		return errs.BadRequest.Wrap(err, errmsg.CheckTimeExpired)
	}

	if err := c.repository.Confirm(*user.Token.Value); err != nil {
		return errs.BadRequest.Wrap(err, errmsg.ConfirmWhileMatching)
	}

	if err := c.email.Send(email.ConfirmData{Email: user.Email}); err != nil {
		return errs.BadRequest.Wrap(err, errmsg.FailedSendMessage)
	}

	return nil
}

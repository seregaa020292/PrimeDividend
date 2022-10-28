package command

import (
	"github.com/google/uuid"

	"primedividend/api/internal/decorators"
	"primedividend/api/internal/modules/auth/repository"
	"primedividend/api/internal/modules/auth/service/email"
	"primedividend/api/pkg/errs"
	"primedividend/api/pkg/errs/errmsg"
)

type (
	ConfirmByToken decorators.CommandHandler[uuid.UUID]
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

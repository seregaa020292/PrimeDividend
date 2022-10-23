package command

import (
	"github.com/google/uuid"

	"primedivident/internal/decorators"
	"primedivident/internal/modules/asset/repository"
	"primedivident/pkg/errs"
	"primedivident/pkg/errs/errmsg"
)

type (
	PayloadRemove struct {
		UserID  uuid.UUID
		AssetID uuid.UUID
	}
	Remove decorators.CommandHandler[PayloadRemove]
)

type remove struct {
	repository repository.Repository
}

func NewRemove(
	repository repository.Repository,
) Remove {
	return remove{
		repository: repository,
	}
}

func (c remove) Exec(cmd PayloadRemove) error {
	exist, err := c.repository.HasByUser(cmd.AssetID, cmd.UserID)
	if err != nil {
		return errs.BadRequest.Wrap(err, errmsg.FailedGetData)
	}

	if !exist {
		return errs.BadRequest.Wrap(err, errmsg.ConfirmWhileMatching)
	}

	if err := c.repository.Remove(cmd.AssetID); err != nil {
		return errs.BadRequest.Wrap(err, errmsg.FailedDeleteData)
	}

	return nil
}

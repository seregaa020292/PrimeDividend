package command

import (
	"github.com/google/uuid"

	"primedividend/api/internal/decorators"
	"primedividend/api/internal/modules/asset/repository"
	"primedividend/api/pkg/errs"
	"primedividend/api/pkg/errs/errmsg"
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

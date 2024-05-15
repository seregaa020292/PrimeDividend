package command

import (
	"context"

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
	Remove decorators.CommandCtxHandler[PayloadRemove]
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

func (c remove) Exec(ctx context.Context, cmd PayloadRemove) error {
	exist, err := c.repository.HasByUser(ctx, cmd.AssetID, cmd.UserID)
	if err != nil {
		return errs.BadRequest.Wrap(err, errmsg.FailedGetData)
	}

	if !exist {
		return errs.BadRequest.Wrap(err, errmsg.ConfirmWhileMatching)
	}

	if err := c.repository.Remove(ctx, cmd.AssetID); err != nil {
		return errs.BadRequest.Wrap(err, errmsg.FailedDeleteData)
	}

	return nil
}

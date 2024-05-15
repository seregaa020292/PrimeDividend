package command

import (
	"context"
	"time"

	"github.com/google/uuid"

	"primedividend/api/internal/decorators"
	"primedividend/api/internal/modules/asset/repository"
	"primedividend/api/pkg/db/transaction"
	"primedividend/api/pkg/errs"
	"primedividend/api/pkg/errs/errmsg"
)

type (
	PayloadUpdate struct {
		UserID  uuid.UUID
		AssetID uuid.UUID

		Quantity   *int32
		Amount     *int32
		NotationAt *time.Time
	}
	Edit decorators.CommandCtxHandler[PayloadUpdate]
)

type edit struct {
	repository repository.Repository
	txManager  transaction.TxManager
}

func NewEdit(
	repository repository.Repository,
	txManager transaction.TxManager,
) Edit {
	return edit{
		repository: repository,
		txManager:  txManager,
	}
}

func (c edit) Exec(ctx context.Context, cmd PayloadUpdate) error {
	exist, err := c.repository.HasByUser(ctx, cmd.AssetID, cmd.UserID)
	if err != nil {
		return errs.BadRequest.Wrap(err, errmsg.FailedGetData)
	}

	if !exist {
		return errs.BadRequest.Wrap(err, errmsg.ConfirmWhileMatching)
	}

	return c.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		if err := c.repository.Update(ctx, cmd.AssetID,
			repository.NewUpdatePatch(
				cmd.Quantity,
				cmd.Amount,
				cmd.NotationAt,
			),
		); err != nil {
			return errs.BadRequest.Wrap(err, errmsg.FailedUpdateData)
		}

		return nil
	})
}

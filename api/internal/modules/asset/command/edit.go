package command

import (
	"time"

	"github.com/google/uuid"

	"primedivident/internal/decorators"
	"primedivident/internal/modules/asset/repository"
	"primedivident/pkg/errs"
	"primedivident/pkg/errs/errmsg"
)

type (
	PayloadUpdate struct {
		UserID  uuid.UUID
		AssetID uuid.UUID

		Quantity   *int32
		Amount     *int32
		NotationAt *time.Time
	}
	Edit decorators.CommandHandler[PayloadUpdate]
)

type edit struct {
	repository repository.Repository
}

func NewEdit(
	repository repository.Repository,
) Edit {
	return edit{
		repository: repository,
	}
}

func (c edit) Exec(cmd PayloadUpdate) error {
	exist, err := c.repository.HasByUser(cmd.AssetID, cmd.UserID)
	if err != nil {
		return errs.BadRequest.Wrap(err, errmsg.FailedGetData)
	}

	if !exist {
		return errs.BadRequest.Wrap(err, errmsg.ConfirmWhileMatching)
	}

	if err := c.repository.Update(
		cmd.AssetID,
		repository.NewUpdatePatch(
			cmd.Quantity,
			cmd.Amount,
			cmd.NotationAt,
		),
	); err != nil {
		return errs.BadRequest.Wrap(err, errmsg.FailedUpdateData)
	}

	return nil
}

package command

import (
	"time"

	"github.com/google/uuid"

	"primedividend/api/internal/decorators"
	"primedividend/api/internal/models/app/public/model"
	"primedividend/api/internal/modules/asset/repository"
	"primedividend/api/pkg/errs"
	"primedividend/api/pkg/errs/errmsg"
)

type (
	PayloadCreate struct {
		UserID      uuid.UUID
		PortfolioID uuid.UUID
		Amount      int32
		Quantity    int32
		MarketID    uuid.UUID
		NotationAt  time.Time
	}
	Create decorators.CommandHandler[PayloadCreate]
)

type create struct {
	repository repository.Repository
}

func NewCreate(
	repository repository.Repository,
) Create {
	return create{
		repository: repository,
	}
}

func (c create) Exec(cmd PayloadCreate) error {
	if err := c.repository.Add(model.Assets{
		Amount:      cmd.Amount,
		Quantity:    cmd.Quantity,
		PortfolioID: cmd.PortfolioID,
		MarketID:    cmd.MarketID,
		NotationAt:  cmd.NotationAt,
	}); err != nil {
		return errs.BadRequest.Wrap(err, errmsg.FailedAddData)
	}

	return nil
}

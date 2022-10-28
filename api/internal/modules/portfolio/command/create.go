package command

import (
	"github.com/google/uuid"

	"primedividend/api/internal/decorators"
	"primedividend/api/internal/models/app/public/model"
	"primedividend/api/internal/modules/portfolio/repository"
	"primedividend/api/pkg/errs"
	"primedividend/api/pkg/errs/errmsg"
)

type (
	PortfolioNew struct {
		Title      string
		UserID     uuid.UUID
		CurrencyID uuid.UUID
	}
	Create decorators.CommandHandler[PortfolioNew]
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

func (c create) Exec(cmd PortfolioNew) error {
	if err := c.repository.Add(model.Portfolios{
		Title:      cmd.Title,
		UserID:     cmd.UserID,
		CurrencyID: cmd.CurrencyID,
		Active:     true,
	}); err != nil {
		return errs.BadRequest.Wrap(err, errmsg.FailedAddData)
	}

	return nil
}

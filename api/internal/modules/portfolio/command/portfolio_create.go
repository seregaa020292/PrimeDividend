package command

import (
	"github.com/google/uuid"

	"primedivident/internal/decorator"
	"primedivident/internal/models/app/public/model"
	"primedivident/internal/modules/portfolio/repository"
	"primedivident/pkg/errs"
	"primedivident/pkg/errs/errmsg"
)

type (
	PortfolioNew struct {
		Title      string
		UserID     uuid.UUID
		CurrencyID uuid.UUID
	}
	PortfolioCreate decorator.CommandHandler[PortfolioNew]
)

type portfolioCreate struct {
	repository repository.Repository
}

func NewPortfolioCreate(
	repository repository.Repository,
) PortfolioCreate {
	return portfolioCreate{
		repository: repository,
	}
}

func (c portfolioCreate) Exec(cmd PortfolioNew) error {
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

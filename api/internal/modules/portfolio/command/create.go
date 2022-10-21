package command

import (
	"github.com/google/uuid"

	"primedivident/internal/decorators"
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

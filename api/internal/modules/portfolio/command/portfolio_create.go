package command

import (
	"github.com/google/uuid"

	"primedivident/internal/decorator"
	"primedivident/internal/modules/portfolio/entity"
	"primedivident/internal/modules/portfolio/repository"
	"primedivident/pkg/errorn"
)

type (
	PortfolioNew struct {
		Title      string
		UserId     uuid.UUID
		CurrencyId uuid.UUID
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
	if err := c.repository.Add(entity.Portfolio{
		Title:      cmd.Title,
		UserID:     cmd.UserId,
		CurrencyID: cmd.CurrencyId,
		Active:     true,
	}); err != nil {
		return errorn.ErrorInsert.Wrap(err)
	}

	return nil
}

package command

import (
	"encoding/json"

	"github.com/google/uuid"

	"primedivident/internal/decorator"
	"primedivident/internal/mistake"
	"primedivident/internal/modules/portfolio/repository"
	"primedivident/internal/services/email"
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
	email      email.FirstTestSend
	repository repository.Repository
}

func NewPortfolioCreate(
	email email.FirstTestSend,
	repository repository.Repository,
) PortfolioCreate {
	return portfolioCreate{
		email:      email,
		repository: repository,
	}
}

func (c portfolioCreate) Exec(cmd PortfolioNew) error {
	portfolioNew, _ := json.MarshalIndent(cmd, "", "\t")

	if err := c.email.Send(portfolioNew); err != nil {
		return mistake.UnknownError(err, "")
	}

	return nil
}

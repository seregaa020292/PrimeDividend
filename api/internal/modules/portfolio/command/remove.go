package command

import (
	"github.com/google/uuid"

	"primedivident/internal/decorators"
	"primedivident/internal/modules/portfolio/repository"
	"primedivident/pkg/errs"
	"primedivident/pkg/errs/errmsg"
)

type (
	PortfolioDelete struct {
		UserID      uuid.UUID
		PortfolioID uuid.UUID
	}
	PortfolioRemove decorators.CommandHandler[PortfolioDelete]
)

type portfolioRemove struct {
	repository repository.Repository
}

func NewPortfolioRemove(
	repository repository.Repository,
) PortfolioRemove {
	return portfolioRemove{
		repository: repository,
	}
}

func (c portfolioRemove) Exec(cmd PortfolioDelete) error {
	if err := c.repository.Remove(cmd.PortfolioID, cmd.UserID); err != nil {
		return errs.BadRequest.Wrap(err, errmsg.FailedDeleteData)
	}

	return nil
}

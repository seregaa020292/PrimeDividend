package command

import (
	"github.com/google/uuid"

	"primedivident/internal/decorators"
	"primedivident/internal/modules/portfolio/dto"
	"primedivident/internal/modules/portfolio/repository"
	"primedivident/pkg/errs"
	"primedivident/pkg/errs/errmsg"
)

type (
	PortfolioUpdate struct {
		UserID      uuid.UUID
		PortfolioID uuid.UUID

		Title      *string
		CurrencyID *uuid.UUID
		Active     *bool
	}
	PortfolioEdit decorators.CommandHandler[PortfolioUpdate]
)

type portfolioEdit struct {
	repository repository.Repository
}

func NewPortfolioEdit(
	repository repository.Repository,
) PortfolioEdit {
	return portfolioEdit{
		repository: repository,
	}
}

func (c portfolioEdit) Exec(cmd PortfolioUpdate) error {
	if err := c.repository.Update(
		cmd.PortfolioID,
		cmd.UserID,
		dto.NewPortfolioVariadic(
			cmd.Title,
			cmd.CurrencyID,
			cmd.Active,
		),
	); err != nil {
		return errs.BadRequest.Wrap(err, errmsg.FailedUpdateData)
	}

	return nil
}

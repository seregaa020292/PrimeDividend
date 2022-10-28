package command

import (
	"github.com/google/uuid"

	"primedividend/api/internal/decorators"
	"primedividend/api/internal/modules/portfolio/repository"
	"primedividend/api/pkg/errs"
	"primedividend/api/pkg/errs/errmsg"
)

type (
	PortfolioUpdate struct {
		UserID      uuid.UUID
		PortfolioID uuid.UUID

		Title      *string
		CurrencyID *uuid.UUID
		Active     *bool
	}
	Edit decorators.CommandHandler[PortfolioUpdate]
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

func (c edit) Exec(cmd PortfolioUpdate) error {
	if err := c.repository.Update(
		cmd.PortfolioID,
		cmd.UserID,
		repository.NewUpdatePatch(
			cmd.Title,
			cmd.CurrencyID,
			cmd.Active,
		),
	); err != nil {
		return errs.BadRequest.Wrap(err, errmsg.FailedUpdateData)
	}

	return nil
}

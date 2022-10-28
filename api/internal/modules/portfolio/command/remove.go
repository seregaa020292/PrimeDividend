package command

import (
	"github.com/google/uuid"

	"primedividend/api/internal/decorators"
	"primedividend/api/internal/modules/portfolio/repository"
	"primedividend/api/pkg/errs"
	"primedividend/api/pkg/errs/errmsg"
)

type (
	PortfolioDelete struct {
		UserID      uuid.UUID
		PortfolioID uuid.UUID
	}
	Remove decorators.CommandHandler[PortfolioDelete]
)

type remove struct {
	repository repository.Repository
}

func NewRemove(
	repository repository.Repository,
) Remove {
	return remove{
		repository: repository,
	}
}

func (c remove) Exec(cmd PortfolioDelete) error {
	if err := c.repository.Remove(cmd.PortfolioID, cmd.UserID); err != nil {
		return errs.BadRequest.Wrap(err, errmsg.FailedDeleteData)
	}

	return nil
}

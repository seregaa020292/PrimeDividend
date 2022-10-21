package query

import (
	"github.com/google/uuid"

	"primedivident/internal/decorators"
	"primedivident/internal/models/app/public/model"
	"primedivident/internal/modules/portfolio/repository"
	"primedivident/pkg/errs"
	"primedivident/pkg/errs/errmsg"
)

type (
	PortfolioId   uuid.UUID
	PortfolioById decorators.QueryHandler[PortfolioId, model.Portfolios]
)

type portfolioById struct {
	repository repository.Repository
}

func NewPortfolioById(
	repository repository.Repository,
) PortfolioById {
	return portfolioById{
		repository: repository,
	}
}

func (q portfolioById) Fetch(portfolioId PortfolioId) (model.Portfolios, error) {
	portfolio, err := q.repository.FindById(uuid.UUID(portfolioId))
	if err != nil {
		return model.Portfolios{}, errs.NotFound.Wrap(err, errmsg.CouldNotBeFound)
	}

	return portfolio, nil
}

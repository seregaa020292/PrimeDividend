package query

import (
	"github.com/google/uuid"

	"primedivident/internal/decorator"
	"primedivident/internal/modules/portfolio/entity"
	"primedivident/internal/modules/portfolio/repository"
	"primedivident/pkg/errs"
	"primedivident/pkg/errs/errmsg"
)

type (
	PortfolioId   uuid.UUID
	PortfolioById decorator.QueryHandler[PortfolioId, entity.Portfolio]
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

func (q portfolioById) Fetch(portfolioId PortfolioId) (entity.Portfolio, error) {
	portfolio, err := q.repository.FindById(uuid.UUID(portfolioId))
	if err != nil {
		return entity.Portfolio{}, errs.NotFound.Wrap(err, errmsg.CouldNotBeFound)
	}

	return portfolio, nil
}

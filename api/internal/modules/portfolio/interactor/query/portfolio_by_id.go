package query

import (
	"github.com/google/uuid"
	"primedivident/internal/decorator"
	"primedivident/internal/mistake"
	"primedivident/internal/modules/portfolio/entity"
	"primedivident/internal/modules/portfolio/repository"
)

type (
	PortfolioId   = uuid.UUID
	PortfolioById = decorator.QueryHandler[PortfolioId, entity.Portfolio]
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
	portfolio, err := q.repository.FindById(portfolioId)
	if err != nil {
		return entity.Portfolio{}, mistake.UnknownError(err, "")
	}

	return portfolio, nil
}

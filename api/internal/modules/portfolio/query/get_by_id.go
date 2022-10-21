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
	ID      uuid.UUID
	GetById decorators.QueryHandler[ID, model.Portfolios]
)

type getById struct {
	repository repository.Repository
}

func NewGetById(
	repository repository.Repository,
) GetById {
	return getById{
		repository: repository,
	}
}

func (q getById) Fetch(portfolioId ID) (model.Portfolios, error) {
	portfolio, err := q.repository.FindById(uuid.UUID(portfolioId))
	if err != nil {
		return model.Portfolios{}, errs.NotFound.Wrap(err, errmsg.CouldNotBeFound)
	}

	return portfolio, nil
}

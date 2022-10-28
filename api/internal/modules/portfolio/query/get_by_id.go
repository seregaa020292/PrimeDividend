package query

import (
	"github.com/google/uuid"

	"primedividend/api/internal/decorators"
	"primedividend/api/internal/models/app/public/model"
	"primedividend/api/internal/modules/portfolio/repository"
	"primedividend/api/pkg/errs"
	"primedividend/api/pkg/errs/errmsg"
)

type (
	ID      = uuid.UUID
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
	portfolio, err := q.repository.FindById(portfolioId)
	if err != nil {
		return model.Portfolios{}, errs.NotFound.Wrap(err, errmsg.CouldNotBeFound)
	}

	return portfolio, nil
}

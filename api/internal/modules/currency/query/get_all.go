package query

import (
	"primedivident/internal/decorators"
	"primedivident/internal/models/app/public/model"
	"primedivident/internal/modules/currency/repository"
	"primedivident/pkg/errs"
	"primedivident/pkg/errs/errmsg"
)

type (
	FilterGetAll struct{}
	GetAll       decorators.QueryHandler[FilterGetAll, []model.Currencies]
)

type getAll struct {
	repository repository.Repository
}

func NewGetAll(
	repository repository.Repository,
) GetAll {
	return getAll{
		repository: repository,
	}
}

func (q getAll) Fetch(filter FilterGetAll) ([]model.Currencies, error) {
	currencies, err := q.repository.GetAll()
	if err != nil {
		return nil, errs.NotFound.Wrap(err, errmsg.FailedGetData)
	}

	return currencies, nil
}

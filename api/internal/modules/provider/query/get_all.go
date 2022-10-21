package query

import (
	"primedivident/internal/decorators"
	"primedivident/internal/models/app/public/model"
	"primedivident/internal/modules/provider/repository"
	"primedivident/pkg/errs"
	"primedivident/pkg/errs/errmsg"
)

type (
	FilterGetAll struct{}
	GetAll       decorators.QueryHandler[FilterGetAll, []model.Providers]
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

func (q getAll) Fetch(filter FilterGetAll) ([]model.Providers, error) {
	providers, err := q.repository.GetAll()
	if err != nil {
		return nil, errs.NotFound.Wrap(err, errmsg.FailedGetData)
	}

	return providers, nil
}

package query

import (
	"primedivident/internal/decorators"
	"primedivident/internal/models/app/public/model"
	"primedivident/internal/modules/instrument/repository"
	"primedivident/pkg/errs"
	"primedivident/pkg/errs/errmsg"
)

type (
	FilterGetAll struct{}
	GetAll       decorators.QueryHandler[FilterGetAll, []model.Instruments]
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

func (q getAll) Fetch(_ FilterGetAll) ([]model.Instruments, error) {
	instruments, err := q.repository.GetAll()
	if err != nil {
		return nil, errs.NotFound.Wrap(err, errmsg.FailedGetData)
	}

	return instruments, nil
}

package query

import (
	"primedividend/api/internal/decorators"
	"primedividend/api/internal/models/app/public/model"
	"primedividend/api/internal/modules/instrument/repository"
	"primedividend/api/pkg/errs"
	"primedividend/api/pkg/errs/errmsg"
)

type (
	PayloadAll struct{}
	GetAll     decorators.QueryHandler[PayloadAll, []model.Instruments]
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

func (q getAll) Fetch(payload PayloadAll) ([]model.Instruments, error) {
	instruments, err := q.repository.GetAll()
	if err != nil {
		return nil, errs.NotFound.Wrap(err, errmsg.FailedGetData)
	}

	return instruments, nil
}

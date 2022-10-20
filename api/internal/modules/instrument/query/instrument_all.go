package query

import (
	"primedivident/internal/decorator"
	"primedivident/internal/models/app/public/model"
	"primedivident/internal/modules/instrument/repository"
	"primedivident/pkg/errs"
	"primedivident/pkg/errs/errmsg"
)

type (
	FilterOrderInstruments struct{}
	InstrumentAll          decorator.QueryHandler[FilterOrderInstruments, []model.Instruments]
)

type instrumentAll struct {
	repository repository.Repository
}

func NewInstrumentAll(
	repository repository.Repository,
) InstrumentAll {
	return instrumentAll{
		repository: repository,
	}
}

func (q instrumentAll) Fetch(_ FilterOrderInstruments) ([]model.Instruments, error) {
	instruments, err := q.repository.GetAll()
	if err != nil {
		return nil, errs.NotFound.Wrap(err, errmsg.FailedGetData)
	}

	return instruments, nil
}

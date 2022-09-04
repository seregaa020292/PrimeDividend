package query

import (
	"primedivident/internal/decorator"
	"primedivident/internal/mistake"
	"primedivident/internal/modules/instrument/entity"
	"primedivident/internal/modules/instrument/repository"
)

type (
	FilterOrderInstruments struct{}
	InstrumentAll          decorator.QueryHandler[FilterOrderInstruments, entity.Instruments]
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

func (q instrumentAll) Fetch(_ FilterOrderInstruments) (entity.Instruments, error) {
	instruments, err := q.repository.GetAll()
	if err != nil {
		return entity.Instruments{}, mistake.UnknownError(err, "")
	}

	return instruments, nil
}

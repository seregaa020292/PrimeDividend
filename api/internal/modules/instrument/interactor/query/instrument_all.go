package query

import (
	"primedivident/internal/decorator"
	"primedivident/internal/mistake"
	"primedivident/internal/modules/instrument/entity"
	"primedivident/internal/modules/instrument/repository"
	"primedivident/internal/services/email"
)

type (
	FilterOrderInstruments struct{}
	InstrumentAll          decorator.QueryHandler[FilterOrderInstruments, entity.Instruments]
)

type instrumentAll struct {
	email      email.FirstTestSend
	repository repository.Repository
}

func NewInstrumentAll(
	email email.FirstTestSend,
	repository repository.Repository,
) InstrumentAll {
	return instrumentAll{
		email:      email,
		repository: repository,
	}
}

func (q instrumentAll) Fetch(_ FilterOrderInstruments) (entity.Instruments, error) {
	instruments, err := q.repository.GetAll()
	if err != nil {
		return entity.Instruments{}, mistake.UnknownError(err, "")
	}

	if err := q.email.Send(); err != nil {
		return entity.Instruments{}, mistake.UnknownError(err, "")
	}

	return instruments, nil
}

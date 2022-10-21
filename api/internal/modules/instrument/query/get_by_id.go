package query

import (
	"github.com/google/uuid"

	"primedivident/internal/decorators"
	"primedivident/internal/models/app/public/model"
	"primedivident/internal/modules/instrument/repository"
	"primedivident/pkg/errs"
	"primedivident/pkg/errs/errmsg"
)

type (
	ID      uuid.UUID
	GetById decorators.QueryHandler[ID, model.Instruments]
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

func (q getById) Fetch(id ID) (model.Instruments, error) {
	instrument, err := q.repository.FindById(uuid.UUID(id))
	if err != nil {
		return model.Instruments{}, errs.NotFound.Wrap(err, errmsg.CouldNotBeFound)
	}

	return instrument, nil
}

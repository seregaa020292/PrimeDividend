package query

import (
	"github.com/google/uuid"

	"primedivident/internal/decorators"
	"primedivident/internal/models/app/public/model"
	"primedivident/internal/modules/currency/repository"
	"primedivident/pkg/errs"
	"primedivident/pkg/errs/errmsg"
)

type (
	ID      = uuid.UUID
	GetById decorators.QueryHandler[ID, model.Currencies]
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

func (q getById) Fetch(id ID) (model.Currencies, error) {
	currency, err := q.repository.FindById(id)
	if err != nil {
		return model.Currencies{}, errs.NotFound.Wrap(err, errmsg.CouldNotBeFound)
	}

	return currency, nil
}

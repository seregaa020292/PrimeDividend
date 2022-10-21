package query

import (
	"github.com/google/uuid"

	"primedivident/internal/decorators"
	"primedivident/internal/models/app/public/model"
	"primedivident/internal/modules/provider/repository"
	"primedivident/pkg/errs"
	"primedivident/pkg/errs/errmsg"
)

type (
	ID      uuid.UUID
	GetById decorators.QueryHandler[ID, model.Providers]
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

func (q getById) Fetch(id ID) (model.Providers, error) {
	provider, err := q.repository.FindById(uuid.UUID(id))
	if err != nil {
		return model.Providers{}, errs.NotFound.Wrap(err, errmsg.CouldNotBeFound)
	}

	return provider, nil
}

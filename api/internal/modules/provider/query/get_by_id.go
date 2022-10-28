package query

import (
	"github.com/google/uuid"

	"primedividend/api/internal/decorators"
	"primedividend/api/internal/models/app/public/model"
	"primedividend/api/internal/modules/provider/repository"
	"primedividend/api/pkg/errs"
	"primedividend/api/pkg/errs/errmsg"
)

type (
	ID      = uuid.UUID
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
	provider, err := q.repository.FindById(id)
	if err != nil {
		return model.Providers{}, errs.NotFound.Wrap(err, errmsg.CouldNotBeFound)
	}

	return provider, nil
}

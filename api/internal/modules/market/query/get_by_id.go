package query

import (
	"github.com/google/uuid"

	"primedivident/internal/decorators"
	"primedivident/internal/models/app/public/model"
	"primedivident/internal/modules/market/repository"
	"primedivident/pkg/errs"
	"primedivident/pkg/errs/errmsg"
)

type (
	ID      uuid.UUID
	GetById decorators.QueryHandler[ID, model.Markets]
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

func (q getById) Fetch(id ID) (model.Markets, error) {
	market, err := q.repository.FindById(uuid.UUID(id))
	if err != nil {
		return model.Markets{}, errs.NotFound.Wrap(err, errmsg.CouldNotBeFound)
	}

	return market, nil
}

package query

import (
	"github.com/google/uuid"

	"primedivident/internal/decorators"
	"primedivident/internal/models"
	"primedivident/internal/models/app/public/model"
	"primedivident/internal/modules/user/repository"
	"primedivident/pkg/errs"
	"primedivident/pkg/errs/errmsg"
)

type (
	ID      = uuid.UUID
	GetById decorators.QueryHandler[ID, model.Users]
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

func (q getById) Fetch(id ID) (model.Users, error) {
	user, err := q.repository.FindById(id, models.ActiveStatus)
	if err != nil {
		return model.Users{}, errs.NotFound.Wrap(err, errmsg.CouldNotBeFound)
	}

	return user, nil
}

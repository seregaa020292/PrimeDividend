package query

import (
	"github.com/google/uuid"

	"primedividend/api/internal/decorators"
	"primedividend/api/internal/models"
	"primedividend/api/internal/models/app/public/model"
	"primedividend/api/internal/modules/user/repository"
	"primedividend/api/pkg/errs"
	"primedividend/api/pkg/errs/errmsg"
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

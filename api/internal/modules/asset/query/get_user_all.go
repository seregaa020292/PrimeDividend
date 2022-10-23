package query

import (
	"github.com/google/uuid"

	"primedivident/internal/decorators"
	"primedivident/internal/models/app/public/model"
	"primedivident/internal/modules/asset/repository"
	"primedivident/pkg/errs"
	"primedivident/pkg/errs/errmsg"
)

type (
	FilterGetUserAll struct {
		UserID      uuid.UUID
		PortfolioID uuid.UUID
	}
	GetUserAll decorators.QueryHandler[FilterGetUserAll, []model.Assets]
)

type getUserAll struct {
	repository repository.Repository
}

func NewGetUserAll(
	repository repository.Repository,
) GetUserAll {
	return getUserAll{
		repository: repository,
	}
}

func (q getUserAll) Fetch(filter FilterGetUserAll) ([]model.Assets, error) {
	assets, err := q.repository.GetUserAll(filter.UserID, filter.PortfolioID)
	if err != nil {
		return nil, errs.NotFound.Wrap(err, errmsg.FailedGetData)
	}

	return assets, nil
}

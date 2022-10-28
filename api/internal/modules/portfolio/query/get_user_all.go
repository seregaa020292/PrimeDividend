package query

import (
	"github.com/google/uuid"

	"primedividend/api/internal/decorators"
	"primedividend/api/internal/models/app/public/model"
	"primedividend/api/internal/modules/portfolio/repository"
	"primedividend/api/pkg/errs"
	"primedividend/api/pkg/errs/errmsg"
)

type (
	PayloadUserAll struct {
		UserID uuid.UUID
		Active *bool
	}
	GetUserAll decorators.QueryHandler[PayloadUserAll, []model.Portfolios]
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

func (q getUserAll) Fetch(payload PayloadUserAll) ([]model.Portfolios, error) {
	portfolios, err := q.repository.GetUserAll(payload.UserID, repository.FilterGetAll{
		Active: payload.Active,
	})
	if err != nil {
		return nil, errs.NotFound.Wrap(err, errmsg.FailedGetData)
	}

	return portfolios, nil
}

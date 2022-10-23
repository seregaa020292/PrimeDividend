package query

import (
	"github.com/google/uuid"

	"primedivident/internal/decorators"
	"primedivident/internal/models/app/public/model"
	"primedivident/internal/modules/portfolio/repository"
	"primedivident/pkg/errs"
	"primedivident/pkg/errs/errmsg"
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

package query

import (
	"context"

	"github.com/google/uuid"

	"primedividend/api/internal/decorators"
	"primedividend/api/internal/models/app/public/model"
	"primedividend/api/internal/modules/asset/repository"
	"primedividend/api/pkg/errs"
	"primedividend/api/pkg/errs/errmsg"
)

type (
	PayloadUserAll struct {
		UserID      uuid.UUID
		PortfolioID uuid.UUID
	}
	GetUserAll decorators.QueryCtxHandler[PayloadUserAll, []model.Assets]
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

func (q getUserAll) Fetch(ctx context.Context, payload PayloadUserAll) ([]model.Assets, error) {
	assets, err := q.repository.GetUserAll(ctx, payload.UserID, payload.PortfolioID)
	if err != nil {
		return nil, errs.NotFound.Wrap(err, errmsg.FailedGetData)
	}

	return assets, nil
}

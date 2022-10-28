package query

import (
	"primedividend/api/internal/decorators"
	"primedividend/api/internal/models/app/public/model"
	"primedividend/api/internal/modules/portfolio/repository"
	"primedividend/api/pkg/errs"
	"primedividend/api/pkg/errs/errmsg"
	"primedividend/api/pkg/paginate/cursor"
)

type (
	PayloadAll struct {
		Limit  *int
		Cursor *string
		Active *bool
	}
	GetAllResult = cursor.PaginateResult[model.Portfolios]
	GetAll       decorators.QueryHandler[PayloadAll, GetAllResult]
)

type getAll struct {
	repository repository.Repository
}

func NewGetAll(
	repository repository.Repository,
) GetAll {
	return getAll{
		repository: repository,
	}
}

func (q getAll) Fetch(payload PayloadAll) (GetAllResult, error) {
	paginateInput, err := cursor.NewPaginateInput(payload.Limit, payload.Cursor)
	if err != nil {
		return GetAllResult{}, errs.BadRequest.Wrap(err, errmsg.UnknownError)
	}

	portfolios, err := q.repository.GetAll(paginateInput, repository.FilterGetAll{
		Active: payload.Active,
	})
	if err != nil {
		return GetAllResult{}, errs.NotFound.Wrap(err, errmsg.FailedGetData)
	}

	return cursor.NewPaginateResult(portfolios, paginateInput), nil
}

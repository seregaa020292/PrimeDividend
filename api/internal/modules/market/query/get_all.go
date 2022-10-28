package query

import (
	"primedividend/api/internal/decorators"
	"primedividend/api/internal/models/app/public/model"
	"primedividend/api/internal/modules/market/repository"
	"primedividend/api/pkg/errs"
	"primedividend/api/pkg/errs/errmsg"
	"primedividend/api/pkg/paginate/cursor"
)

type (
	PayloadAll struct {
		Limit  *int
		Cursor *string
	}
	GetAllResult = cursor.PaginateResult[model.Markets]
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

	markets, err := q.repository.GetAll(paginateInput)
	if err != nil {
		return GetAllResult{}, errs.NotFound.Wrap(err, errmsg.FailedGetData)
	}

	return cursor.NewPaginateResult(markets, paginateInput), nil
}

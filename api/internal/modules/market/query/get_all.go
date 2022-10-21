package query

import (
	"primedivident/internal/decorators"
	"primedivident/internal/models/app/public/model"
	"primedivident/internal/modules/market/repository"
	"primedivident/pkg/errs"
	"primedivident/pkg/errs/errmsg"
	"primedivident/pkg/paginate/cursor"
)

type (
	FilterGetAll struct {
		Limit  *int
		Cursor *string
	}
	GetAllResult = cursor.PaginateResult[model.Markets]
	GetAll       decorators.QueryHandler[FilterGetAll, GetAllResult]
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

func (q getAll) Fetch(filter FilterGetAll) (GetAllResult, error) {
	paginateInput, err := cursor.NewPaginateInput(filter.Limit, filter.Cursor)
	if err != nil {
		return GetAllResult{}, errs.BadRequest.Wrap(err, errmsg.UnknownError)
	}

	markets, err := q.repository.GetAll(paginateInput)
	if err != nil {
		return GetAllResult{}, errs.NotFound.Wrap(err, errmsg.FailedGetData)
	}

	return cursor.NewPaginateResult(markets, paginateInput), nil
}

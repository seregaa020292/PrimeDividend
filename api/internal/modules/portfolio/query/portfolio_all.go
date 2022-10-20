package query

import (
	"primedivident/internal/decorator"
	"primedivident/internal/models/app/public/model"
	"primedivident/internal/modules/portfolio/repository"
	"primedivident/pkg/errs"
	"primedivident/pkg/errs/errmsg"
	"primedivident/pkg/utils/paginate/cursor"
)

type (
	PortfoliosInput struct {
		Limit  *int
		Cursor *string
		Active bool
	}
	PortfoliosResult = cursor.PaginateResult[model.Portfolios]
	PortfolioAll     decorator.QueryHandler[PortfoliosInput, PortfoliosResult]
)

type portfolioAll struct {
	repository repository.Repository
}

func NewPortfolioAll(
	repository repository.Repository,
) PortfolioAll {
	return portfolioAll{
		repository: repository,
	}
}

func (q portfolioAll) Fetch(filter PortfoliosInput) (PortfoliosResult, error) {
	paginateInput, err := cursor.NewPaginateInput(filter.Limit, filter.Cursor)
	if err != nil {
		return PortfoliosResult{}, errs.BadRequest.Wrap(err, errmsg.UnknownError)
	}

	portfolios, err := q.repository.GetAll(paginateInput, model.Portfolios{
		Active: filter.Active,
	})
	if err != nil {
		return PortfoliosResult{}, errs.NotFound.Wrap(err, errmsg.FailedGetData)
	}

	return cursor.NewPaginateResult(portfolios, paginateInput), nil
}

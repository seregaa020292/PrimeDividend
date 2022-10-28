package query

import (
	"primedividend/api/internal/decorators"
	"primedividend/api/internal/models/app/public/model"
	"primedividend/api/internal/modules/market/repository"
	"primedividend/api/pkg/errs"
	"primedividend/api/pkg/errs/errmsg"
)

type (
	Ticker      = string
	GetByTicker decorators.QueryHandler[Ticker, model.Markets]
)

type getByTicker struct {
	repository repository.Repository
}

func NewGetByTicker(
	repository repository.Repository,
) GetByTicker {
	return getByTicker{
		repository: repository,
	}
}

func (q getByTicker) Fetch(ticker Ticker) (model.Markets, error) {
	market, err := q.repository.FindByTicker(ticker)
	if err != nil {
		return model.Markets{}, errs.NotFound.Wrap(err, errmsg.CouldNotBeFound)
	}

	return market, nil
}

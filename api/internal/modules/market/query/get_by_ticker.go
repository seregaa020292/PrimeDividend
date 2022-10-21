package query

import (
	"primedivident/internal/decorators"
	"primedivident/internal/models/app/public/model"
	"primedivident/internal/modules/market/repository"
	"primedivident/pkg/errs"
	"primedivident/pkg/errs/errmsg"
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

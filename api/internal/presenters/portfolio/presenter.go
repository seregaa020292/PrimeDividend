package portfolio

import (
	"primedivident/internal/infrastructure/server/openapi"
	"primedivident/internal/modules/portfolio/entity"
)

type Presenter interface {
	GetPortfolio(portfolio entity.Portfolio) openapi.Portfolio
}

type present struct {
}

func NewPresenter() Presenter {
	return present{}
}

func (p present) GetPortfolio(portfolio entity.Portfolio) openapi.Portfolio {
	return openapi.Portfolio{
		Id:        portfolio.ID,
		CreatedAt: portfolio.CreatedAt,
	}
}

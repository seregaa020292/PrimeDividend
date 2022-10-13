package portfolio

import (
	"primedivident/internal/infrastructure/server/openapi"
	"primedivident/internal/models/app/public/model"
)

type Presenter interface {
	GetPortfolio(portfolio model.Portfolios) openapi.Portfolio
}

type present struct {
}

func NewPresenter() Presenter {
	return present{}
}

func (p present) GetPortfolio(portfolio model.Portfolios) openapi.Portfolio {
	return openapi.Portfolio{
		Id:         portfolio.ID,
		Title:      portfolio.Title,
		Active:     portfolio.Active,
		CurrencyId: portfolio.CurrencyID,
		UserId:     portfolio.UserID,
		CreatedAt:  portfolio.CreatedAt,
		UpdatedAt:  portfolio.UpdatedAt,
	}
}

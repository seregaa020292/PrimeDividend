package portfolio

import (
	"primedivident/internal/infrastructure/server/openapi"
	"primedivident/internal/models/app/public/model"
	"primedivident/internal/modules/portfolio/query"
)

type Presenter interface {
	GetPortfolio(portfolio model.Portfolios) openapi.Portfolio
	GetPortfolios(portfolios []model.Portfolios) openapi.Portfolios
	GetRecordsMeta(result query.PortfoliosResult) (openapi.Portfolios, openapi.Meta)
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

func (p present) GetPortfolios(portfolios []model.Portfolios) openapi.Portfolios {
	result := make(openapi.Portfolios, len(portfolios))

	for i, item := range portfolios {
		result[i] = openapi.Portfolio{
			Id:         item.ID,
			Title:      item.Title,
			CurrencyId: item.CurrencyID,
			UserId:     item.UserID,
			CreatedAt:  item.CreatedAt,
			UpdatedAt:  item.UpdatedAt,
		}
	}

	return result
}

func (p present) GetRecordsMeta(result query.PortfoliosResult) (openapi.Portfolios, openapi.Meta) {
	return p.GetPortfolios(result.Records), openapi.Meta{
		Pagination: openapi.PagingCursor{
			Count:      result.Length,
			Limit:      result.Limit,
			CursorNext: result.CursorNext,
			CursorPrev: result.CursorPrev,
		},
	}
}

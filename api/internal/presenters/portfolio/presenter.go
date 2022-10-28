package portfolio

import (
	"primedividend/api/internal/infrastructure/server/openapi"
	"primedividend/api/internal/models/app/public/model"
	"primedividend/api/internal/modules/portfolio/query"
)

type Presenter interface {
	GetOne(portfolio model.Portfolios) openapi.Portfolio
	GetAll(portfolios []model.Portfolios) openapi.Portfolios
	GetAllMeta(result query.GetAllResult) (openapi.Portfolios, openapi.Meta)
}

type present struct{}

func NewPresenter() Presenter {
	return present{}
}

func (p present) GetOne(item model.Portfolios) openapi.Portfolio {
	return openapi.Portfolio{
		Id:         item.ID,
		Title:      item.Title,
		CurrencyId: item.CurrencyID,
		UserId:     item.UserID,
		CreatedAt:  item.CreatedAt,
		UpdatedAt:  item.UpdatedAt,
	}
}

func (p present) GetAll(items []model.Portfolios) openapi.Portfolios {
	result := make(openapi.Portfolios, len(items))

	for i, item := range items {
		result[i] = p.GetOne(item)
	}

	return result
}

func (p present) GetAllMeta(result query.GetAllResult) (openapi.Portfolios, openapi.Meta) {
	return p.GetAll(result.Records), openapi.Meta{
		Pagination: openapi.PagingCursor{
			Count:      result.Length,
			Limit:      result.Limit,
			CursorNext: result.CursorNext,
			CursorPrev: result.CursorPrev,
		},
	}
}

package market

import (
	"primedivident/internal/infrastructure/server/openapi"
	"primedivident/internal/models/app/public/model"
	"primedivident/internal/modules/market/query"
)

type Presenter interface {
	GetOne(item model.Markets) openapi.Market
	GetAll(items []model.Markets) openapi.Markets
	GetAllMeta(result query.GetAllResult) (openapi.Markets, openapi.Meta)
}

type present struct{}

func NewPresenter() Presenter {
	return present{}
}

func (p present) GetOne(item model.Markets) openapi.Market {
	return openapi.Market{
		Id:           item.ID,
		Title:        item.Title,
		Ticker:       item.Ticker,
		Content:      item.Content,
		ImageUrl:     item.ImageURL,
		CurrencyId:   item.CurrencyID,
		InstrumentId: item.InstrumentID,
		CreatedAt:    item.CreatedAt,
		UpdatedAt:    item.UpdatedAt,
	}
}

func (p present) GetAll(items []model.Markets) openapi.Markets {
	result := make(openapi.Markets, len(items))

	for i, item := range items {
		result[i] = p.GetOne(item)
	}

	return result
}

func (p present) GetAllMeta(result query.GetAllResult) (openapi.Markets, openapi.Meta) {
	return p.GetAll(result.Records), openapi.Meta{
		Pagination: openapi.PagingCursor{
			Count:      result.Length,
			Limit:      result.Limit,
			CursorNext: result.CursorNext,
			CursorPrev: result.CursorPrev,
		},
	}
}

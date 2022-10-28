package asset

import (
	"time"

	"primedividend/api/internal/infrastructure/server/openapi"
	"primedividend/api/internal/models/app/public/model"
)

type Presenter interface {
	GetOne(portfolio model.Assets) openapi.Asset
	GetAll(portfolios []model.Assets) openapi.Assets
}

type present struct{}

func NewPresenter() Presenter {
	return present{}
}

func (p present) GetOne(item model.Assets) openapi.Asset {
	return openapi.Asset{
		Id:          item.ID,
		Amount:      item.Amount,
		Quantity:    item.Quantity,
		MarketId:    item.MarketID,
		PortfolioId: item.PortfolioID,
		NotationAt:  time.Time{},
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
	}
}

func (p present) GetAll(items []model.Assets) openapi.Assets {
	result := make(openapi.Assets, len(items))

	for i, item := range items {
		result[i] = p.GetOne(item)
	}

	return result
}

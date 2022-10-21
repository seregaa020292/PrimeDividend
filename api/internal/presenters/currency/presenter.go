package currency

import (
	"primedivident/internal/infrastructure/server/openapi"
	"primedivident/internal/models/app/public/model"
)

type Presenter interface {
	GetOne(item model.Currencies) openapi.Currency
	GetAll(items []model.Currencies) openapi.Currencies
}

type present struct{}

func NewPresenter() Presenter {
	return present{}
}

func (p present) GetOne(item model.Currencies) openapi.Currency {
	return openapi.Currency{
		Id:          item.ID,
		Title:       item.Title,
		Description: item.Description,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
	}
}

func (p present) GetAll(items []model.Currencies) openapi.Currencies {
	result := make(openapi.Currencies, len(items))

	for i, item := range items {
		result[i] = p.GetOne(item)
	}

	return result
}

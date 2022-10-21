package provider

import (
	"primedivident/internal/infrastructure/server/openapi"
	"primedivident/internal/models/app/public/model"
)

type Presenter interface {
	GetOne(item model.Providers) openapi.Provider
	GetAll(items []model.Providers) openapi.Providers
}

type present struct{}

func NewPresenter() Presenter {
	return present{}
}

func (p present) GetOne(item model.Providers) openapi.Provider {
	return openapi.Provider{
		Id:          item.ID,
		Title:       item.Title,
		Description: item.Description,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
	}
}

func (p present) GetAll(items []model.Providers) openapi.Providers {
	result := make(openapi.Providers, len(items))

	for i, item := range items {
		result[i] = p.GetOne(item)
	}

	return result
}

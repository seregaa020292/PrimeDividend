package instrument

import (
	"primedividend/api/internal/infrastructure/server/openapi"
	"primedividend/api/internal/models/app/public/model"
)

type Presenter interface {
	GetOne(item model.Instruments) openapi.Instrument
	GetAll(items []model.Instruments) openapi.Instruments
}

type present struct{}

func NewPresenter() Presenter {
	return present{}
}

func (p present) GetOne(item model.Instruments) openapi.Instrument {
	return openapi.Instrument{
		Id:          item.ID,
		Title:       item.Title,
		Description: item.Description,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
	}
}

func (p present) GetAll(items []model.Instruments) openapi.Instruments {
	result := make(openapi.Instruments, len(items))

	for i, item := range items {
		result[i] = p.GetOne(item)
	}

	return result
}

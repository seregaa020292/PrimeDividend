package instrument

import (
	"primedivident/internal/infrastructure/http/openapi"
	"primedivident/internal/models/app/public/model"
)

type Presenter interface {
	GetInstruments(instruments []model.Instruments) openapi.Instruments
}

type present struct {
}

func NewPresenter() Presenter {
	return present{}
}

func (p present) GetInstruments(instruments []model.Instruments) openapi.Instruments {
	result := make(openapi.Instruments, len(instruments))

	for i, item := range instruments {
		result[i] = openapi.Instrument{
			Id:          item.ID,
			Title:       item.Title,
			Description: item.Description,
		}
	}

	return result
}

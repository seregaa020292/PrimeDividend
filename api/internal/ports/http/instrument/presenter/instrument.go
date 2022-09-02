package presenter

import (
	"primedivident/internal/modules/instrument/entity"
	"primedivident/internal/ports/http/instrument"
)

type Presenter interface {
	GetAll(instruments entity.Instruments) instrument.Instruments
}

type Present struct {
}

func NewPresenter() Presenter {
	return Present{}
}

func (p Present) GetAll(instruments entity.Instruments) instrument.Instruments {
	result := make(instrument.Instruments, len(instruments))

	for i, item := range instruments {
		result[i] = instrument.Instrument{
			Id:          item.ID,
			Title:       item.Title,
			Description: item.Description,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		}
	}

	return result
}

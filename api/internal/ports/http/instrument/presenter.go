package instrument

import (
	"primedivident/internal/modules/instrument/entity"
)

func presenterGetAll(instruments entity.Instruments) Instruments {
	result := make(Instruments, len(instruments))

	for i, item := range instruments {
		result[i] = Instrument{
			Id:          item.ID,
			Title:       item.Title,
			Description: item.Description,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		}
	}

	return result
}

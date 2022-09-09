package instrument

import (
	"primedivident/internal/infrastructure/http/openapi"
	"primedivident/internal/modules/instrument/entity"
)

func presenterGetAll(instruments entity.Instruments) openapi.Instruments {
	result := make(openapi.Instruments, len(instruments))

	for i, item := range instruments {
		result[i] = openapi.Instrument{
			Id:          item.ID,
			Title:       item.Title,
			Description: item.Description,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		}
	}

	return result
}

package instrument

import (
	"net/http"

	"primedivident/internal/infrastructure/http/openapi"
	"primedivident/internal/modules/instrument/entity"
	"primedivident/internal/modules/instrument/interactor/query"
)

func (h HandlerInstrument) GetInstruments(w http.ResponseWriter, r *http.Request) {
	respond := h.responder.Http(w, r)

	instruments, err := h.queryInstrumentAll.Fetch(query.FilterOrderInstruments{})
	if err != nil {
		respond.Err(err)
		return
	}

	respond.Json(http.StatusOK, presenterGetAll(instruments))
}

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

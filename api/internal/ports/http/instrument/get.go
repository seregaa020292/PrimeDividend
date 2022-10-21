package instrument

import (
	"net/http"

	"primedivident/internal/infrastructure/server/openapi"
	"primedivident/internal/modules/instrument/query"
)

func (h HandlerInstrument) GetInstrument(w http.ResponseWriter, r *http.Request, instrumentId openapi.InstrumentId) {
	respond := h.responder.Http(w, r)

	instrument, err := h.queryGetById.Fetch(query.ID(instrumentId))
	if err != nil {
		respond.Err(err)
		return
	}

	respond.Json(http.StatusOK, h.presenter.GetOne(instrument))
}

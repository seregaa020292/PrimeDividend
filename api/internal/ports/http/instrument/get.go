package instrument

import (
	"net/http"

	"primedividend/api/internal/infrastructure/server/openapi"
)

func (h HandlerInstrument) GetInstrument(w http.ResponseWriter, r *http.Request, instrumentId openapi.InstrumentId) {
	respond := h.responder.Http(w, r)

	instrument, err := h.queryGetById.Fetch(instrumentId)
	if err != nil {
		respond.Err(err)
		return
	}

	respond.Json(http.StatusOK, h.presenter.GetOne(instrument))
}

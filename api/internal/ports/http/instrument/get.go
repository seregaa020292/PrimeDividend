package instrument

import (
	"net/http"

	"primedivident/internal/infrastructure/server/openapi"
)

func (h HandlerInstrument) GetInstrument(w http.ResponseWriter, r *http.Request, instrumentId openapi.InstrumentId) {
	respond := h.responder.Http(w, r)

	respond.Json(http.StatusOK, nil)
}

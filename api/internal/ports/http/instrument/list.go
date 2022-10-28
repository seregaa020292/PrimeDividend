package instrument

import (
	"net/http"

	"primedividend/api/internal/modules/instrument/query"
)

func (h HandlerInstrument) GetInstruments(w http.ResponseWriter, r *http.Request) {
	respond := h.responder.Http(w, r)

	instruments, err := h.queryGetAll.Fetch(query.PayloadAll{})
	if err != nil {
		respond.Err(err)
		return
	}

	respond.Json(http.StatusOK, h.presenter.GetAll(instruments))
}

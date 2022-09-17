package instrument

import (
	"net/http"

	"primedivident/internal/modules/instrument/query"
)

func (h HandlerInstrument) GetInstruments(w http.ResponseWriter, r *http.Request) {
	respond := h.responder.Http(w, r)

	instruments, err := h.queryInstrumentAll.Fetch(query.FilterOrderInstruments{})
	if err != nil {
		respond.Err(err)
		return
	}

	respond.Json(http.StatusOK, h.presenter.GetInstruments(instruments))
}

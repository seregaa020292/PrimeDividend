package currency

import (
	"net/http"

	"primedivident/internal/modules/currency/query"
)

func (h HandlerCurrency) GetCurrencies(w http.ResponseWriter, r *http.Request) {
	respond := h.responder.Http(w, r)

	instruments, err := h.queryGetAll.Fetch(query.FilterGetAll{})
	if err != nil {
		respond.Err(err)
		return
	}

	respond.Json(http.StatusOK, h.presenter.GetAll(instruments))
}

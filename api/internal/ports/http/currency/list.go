package currency

import "net/http"

func (h HandlerCurrency) GetCurrencies(w http.ResponseWriter, r *http.Request) {
	respond := h.responder.Http(w, r)

	respond.Json(http.StatusOK, []any{})
}

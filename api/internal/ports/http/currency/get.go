package currency

import (
	"net/http"

	"primedividend/api/internal/infrastructure/server/openapi"
)

func (h HandlerCurrency) GetCurrency(w http.ResponseWriter, r *http.Request, currencyId openapi.CurrencyId) {
	respond := h.responder.Http(w, r)

	currency, err := h.queryGetById.Fetch(currencyId)
	if err != nil {
		respond.Err(err)
		return
	}

	respond.Json(http.StatusOK, h.presenter.GetOne(currency))
}

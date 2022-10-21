package currency

import (
	"net/http"

	"primedivident/internal/infrastructure/server/openapi"
	"primedivident/internal/modules/currency/query"
)

func (h HandlerCurrency) GetCurrency(w http.ResponseWriter, r *http.Request, currencyId openapi.CurrencyId) {
	respond := h.responder.Http(w, r)

	currency, err := h.queryGetById.Fetch(query.ID(currencyId))
	if err != nil {
		respond.Err(err)
		return
	}

	respond.Json(http.StatusOK, h.presenter.GetOne(currency))
}

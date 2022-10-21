package currency

import (
	"net/http"

	"primedivident/internal/infrastructure/server/openapi"
)

func (h HandlerCurrency) GetCurrency(w http.ResponseWriter, r *http.Request, currencyId openapi.CurrencyId) {
	respond := h.responder.Http(w, r)

	respond.Json(http.StatusOK, nil)
}

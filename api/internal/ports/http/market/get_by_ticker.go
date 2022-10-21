package market

import (
	"net/http"

	"primedivident/internal/infrastructure/server/openapi"
)

func (h HandlerMarket) GetMarketByTicker(w http.ResponseWriter, r *http.Request, ticker openapi.Ticker) {
	respond := h.responder.Http(w, r)

	market, err := h.queryGetByTicker.Fetch(ticker)
	if err != nil {
		respond.Err(err)
		return
	}

	respond.Json(http.StatusOK, h.presenter.GetOne(market))
}

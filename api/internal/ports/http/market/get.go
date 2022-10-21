package market

import (
	"net/http"

	"primedivident/internal/infrastructure/server/openapi"
)

func (h HandlerMarket) GetMarket(w http.ResponseWriter, r *http.Request, marketId openapi.MarketId) {
	respond := h.responder.Http(w, r)

	market, err := h.queryGetById.Fetch(marketId)
	if err != nil {
		respond.Err(err)
		return
	}

	respond.Json(http.StatusOK, h.presenter.GetOne(market))
}

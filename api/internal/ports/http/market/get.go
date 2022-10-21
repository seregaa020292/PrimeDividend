package market

import (
	"net/http"

	"primedivident/internal/infrastructure/server/openapi"
	"primedivident/internal/modules/market/query"
)

func (h HandlerMarket) GetMarket(w http.ResponseWriter, r *http.Request, marketId openapi.MarketId) {
	respond := h.responder.Http(w, r)

	instrument, err := h.queryGetById.Fetch(query.ID(marketId))
	if err != nil {
		respond.Err(err)
		return
	}

	respond.Json(http.StatusOK, h.presenter.GetOne(instrument))
}

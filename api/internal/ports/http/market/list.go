package market

import (
	"net/http"

	"primedivident/internal/infrastructure/server/openapi"
	"primedivident/internal/modules/market/query"
)

func (h HandlerMarket) GetMarkets(w http.ResponseWriter, r *http.Request, params openapi.GetMarketsParams) {
	respond := h.responder.Http(w, r)

	markets, err := h.queryGetAll.Fetch(query.PayloadAll{
		Limit:  params.Limit,
		Cursor: params.Cursor,
	})
	if err != nil {
		respond.Err(err)
		return
	}

	data, meta := h.presenter.GetAllMeta(markets)

	respond.Json(http.StatusOK, data, meta)
}

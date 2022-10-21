package market

import (
	"net/http"

	"primedivident/internal/infrastructure/server/openapi"
	"primedivident/internal/modules/market/query"
)

func (h HandlerMarket) GetMarkets(w http.ResponseWriter, r *http.Request, params openapi.GetMarketsParams) {
	respond := h.responder.Http(w, r)

	instruments, err := h.queryGetAll.Fetch(query.FilterGetAll{
		Limit:  params.Limit,
		Cursor: params.Cursor,
	})
	if err != nil {
		respond.Err(err)
		return
	}

	data, meta := h.presenter.GetAllMeta(instruments)

	respond.Json(http.StatusOK, data, meta)
}
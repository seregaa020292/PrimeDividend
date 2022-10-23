package portfolio

import (
	"net/http"

	"primedivident/internal/infrastructure/server/openapi"
	"primedivident/internal/modules/portfolio/query"
	"primedivident/pkg/utils/gog"
)

func (h HandlerPortfolio) GetPortfolios(w http.ResponseWriter, r *http.Request, params openapi.GetPortfoliosParams) {
	respond := h.responder.Http(w, r)

	portfolios, err := h.queryGetAll.Fetch(query.PayloadAll{
		Limit:  params.Limit,
		Cursor: params.Cursor,
		Active: gog.Ptr(true),
	})
	if err != nil {
		respond.Err(err)
		return
	}

	data, meta := h.presenter.GetAllMeta(portfolios)

	respond.Json(http.StatusOK, data, meta)
}

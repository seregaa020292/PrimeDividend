package portfolio

import (
	"net/http"

	"primedivident/internal/infrastructure/server/openapi"
	"primedivident/internal/modules/portfolio/query"
)

func (h HandlerPortfolio) GetPortfolios(w http.ResponseWriter, r *http.Request, params openapi.GetPortfoliosParams) {
	respond := h.responder.Http(w, r)

	portfolios, err := h.queryPortfolioAll.Fetch(query.PortfoliosInput{
		Limit:  params.Limit,
		Cursor: params.Cursor,
		Active: true,
	})
	if err != nil {
		respond.Err(err)
		return
	}

	data, meta := h.presenter.GetRecordsMeta(portfolios)

	respond.Json(http.StatusOK, data, meta)
}

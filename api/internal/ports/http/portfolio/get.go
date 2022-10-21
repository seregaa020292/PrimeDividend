package portfolio

import (
	"net/http"

	"primedivident/internal/infrastructure/server/openapi"
	"primedivident/internal/modules/portfolio/query"
)

func (h HandlerPortfolio) GetPortfolio(w http.ResponseWriter, r *http.Request, portfolioId openapi.PortfolioId) {
	respond := h.responder.Http(w, r)

	portfolio, err := h.queryGetById.Fetch(query.ID(portfolioId))
	if err != nil {
		respond.Err(err)
		return
	}

	respond.Json(http.StatusOK, h.presenter.GetOne(portfolio))
}

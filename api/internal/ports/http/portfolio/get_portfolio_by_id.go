package portfolio

import (
	"net/http"

	"primedivident/internal/infrastructure/http/openapi"
	"primedivident/internal/modules/portfolio/query"
)

func (h HandlerPortfolio) GetPortfolioById(w http.ResponseWriter, r *http.Request, portfolioId openapi.PortfolioId) {
	respond := h.responder.Http(w, r)

	portfolio, err := h.queryPortfolioById.Fetch(query.PortfolioId(portfolioId))
	if err != nil {
		respond.Err(err)
		return
	}

	respond.Json(http.StatusOK, h.presenter.GetPortfolio(portfolio))
}

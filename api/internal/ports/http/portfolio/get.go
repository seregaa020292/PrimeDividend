package portfolio

import (
	"net/http"

	"primedivident/internal/infrastructure/server/openapi"
)

func (h HandlerPortfolio) GetPortfolio(w http.ResponseWriter, r *http.Request, portfolioId openapi.PortfolioId) {
	respond := h.responder.Http(w, r)

	portfolio, err := h.queryGetById.Fetch(portfolioId)
	if err != nil {
		respond.Err(err)
		return
	}

	respond.Json(http.StatusOK, h.presenter.GetOne(portfolio))
}

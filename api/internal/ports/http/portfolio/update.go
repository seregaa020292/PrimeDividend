package portfolio

import (
	"net/http"

	"primedividend/api/internal/infrastructure/server/middlewares/helper"
	"primedividend/api/internal/infrastructure/server/openapi"
	"primedividend/api/internal/modules/portfolio/command"
)

func (h HandlerPortfolio) UpdatePortfolio(w http.ResponseWriter, r *http.Request, portfolioId openapi.PortfolioId) {
	respond := h.responder.Http(w, r)

	user, err := helper.UserFromCtx(r.Context())
	if err != nil {
		respond.Err(err)
		return
	}

	var portfolio openapi.PortfolioUpdate

	if err := respond.DecodeValidate(&portfolio); err != nil {
		respond.Err(err)
		return
	}

	if err := h.cmdEdit.Exec(command.PortfolioUpdate{
		UserID:      user.ID,
		PortfolioID: portfolioId,
		Title:       portfolio.Title,
		CurrencyID:  portfolio.CurrencyId,
		Active:      portfolio.Active,
	}); err != nil {
		respond.Err(err)
		return
	}

	respond.WriteHeader(http.StatusOK)
}

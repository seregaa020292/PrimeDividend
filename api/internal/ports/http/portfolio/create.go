package portfolio

import (
	"net/http"

	"primedividend/api/internal/infrastructure/server/middlewares/helper"
	"primedividend/api/internal/infrastructure/server/openapi"
	"primedividend/api/internal/modules/portfolio/command"
)

func (h HandlerPortfolio) CreatePortfolio(w http.ResponseWriter, r *http.Request) {
	respond := h.responder.Http(w, r)

	user, err := helper.UserFromCtx(r.Context())
	if err != nil {
		respond.Err(err)
		return
	}

	var portfolio openapi.PortfolioAdd

	if err := respond.DecodeValidate(&portfolio); err != nil {
		respond.Err(err)
		return
	}

	if err := h.cmdCreate.Exec(command.PortfolioNew{
		Title:      portfolio.Title,
		UserID:     user.ID,
		CurrencyID: portfolio.CurrencyId,
	}); err != nil {
		respond.Err(err)
		return
	}

	respond.WriteHeader(http.StatusNoContent)
}

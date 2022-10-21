package portfolio

import (
	"net/http"

	"primedivident/internal/infrastructure/server/middlewares/helper"
	"primedivident/internal/infrastructure/server/openapi"
	"primedivident/internal/modules/portfolio/command"
)

func (h HandlerPortfolio) RemovePortfolio(w http.ResponseWriter, r *http.Request, portfolioId openapi.PortfolioId) {
	respond := h.responder.Http(w, r)

	user, err := helper.UserFromCtx(r.Context())
	if err != nil {
		respond.Err(err)
		return
	}

	if err := h.cmdRemove.Exec(command.PortfolioDelete{
		UserID:      user.ID,
		PortfolioID: portfolioId,
	}); err != nil {
		respond.Err(err)
		return
	}

	respond.WriteHeader(http.StatusOK)
}

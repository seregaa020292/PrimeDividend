package portfolio

import (
	"net/http"

	"primedivident/internal/infrastructure/http/openapi"
	"primedivident/internal/modules/portfolio/interactor/command"
)

func (h HandlerPortfolio) CreatePortfolio(w http.ResponseWriter, r *http.Request) {
	respond := h.responder.Http(w, r)

	portfolio := openapi.PortfolioUpdate{}
	if err := respond.DecodeValidate(&portfolio); err != nil {
		respond.Err(err)
		return
	}

	if err := h.cmdPortfolioCreate.Exec(command.PortfolioNew{
		Title:      portfolio.Title,
		UserId:     portfolio.UserId,
		CurrencyId: portfolio.CurrencyId,
	}); err != nil {
		respond.Err(err)
		return
	}

	respond.NoContent()
}

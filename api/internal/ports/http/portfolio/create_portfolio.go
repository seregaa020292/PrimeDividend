package portfolio

import (
	"net/http"
	"primedivident/internal/infrastructure/http/openapi"
	"primedivident/internal/modules/portfolio/interactor/command"
	"primedivident/pkg/response"
)

func (h HandlerPortfolio) CreatePortfolio(w http.ResponseWriter, r *http.Request) {
	respond := response.New(w, r)

	portfolio := openapi.PortfolioUpdate{}
	if err := respond.Decode(&portfolio); err != nil {
		respond.Err(err)
		return
	}

	if err := h.validator.Struct(portfolio); err != nil {
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

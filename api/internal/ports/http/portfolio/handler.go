package portfolio

import (
	"net/http"

	"primedivident/internal/infrastructure/http/openapi"
	"primedivident/internal/modules/portfolio/interactor/command"
	"primedivident/internal/modules/portfolio/interactor/query"
	"primedivident/pkg/response"
	"primedivident/pkg/validator"
)

type HandlerPortfolio struct {
	queryPortfolioById query.PortfolioById
	cmdPortfolioCreate command.PortfolioCreate
}

func NewHandler(
	validator validator.Validator,
	queryPortfolioById query.PortfolioById,
	cmdPortfolioCreate command.PortfolioCreate,
) HandlerPortfolio {
	return HandlerPortfolio{
		queryPortfolioById: queryPortfolioById,
		cmdPortfolioCreate: cmdPortfolioCreate,
	}
}

func (h HandlerPortfolio) GetPortfolioById(w http.ResponseWriter, r *http.Request, portfolioId openapi.PortfolioId) {
	respond := response.New(w, r)

	portfolio, err := h.queryPortfolioById.Fetch(query.PortfolioId(portfolioId))
	if err != nil {
		respond.Err(err)
		return
	}

	respond.Json(http.StatusOK, presenterPortfolio(portfolio))
}

func (h HandlerPortfolio) CreatePortfolio(w http.ResponseWriter, r *http.Request) {
	respond := response.New(w, r)

	//portfolio := r.Context().Value(portfolioUpdateKey).(openapi.PortfolioUpdate)

	portfolio := openapi.PortfolioUpdate{}
	if err := respond.Decode(&portfolio); err != nil {
		respond.Err(err)
		return
	}

	if err := validator.GetValidator().Struct(portfolio); err != nil {
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

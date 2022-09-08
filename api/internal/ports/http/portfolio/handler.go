package portfolio

import (
	"net/http"

	"primedivident/internal/modules/portfolio/interactor/command"
	"primedivident/internal/modules/portfolio/interactor/query"
	"primedivident/pkg/response"
	"primedivident/pkg/validator"
)

//var _ ServerInterface = (*handler)(nil)

type handler struct {
	validator          validator.Validator
	queryPortfolioById query.PortfolioById
	cmdPortfolioCreate command.PortfolioCreate
}

func NewHandler(
	validator validator.Validator,
	queryPortfolioById query.PortfolioById,
	cmdPortfolioCreate command.PortfolioCreate,
) ServerInterface {
	return handler{
		validator:          validator,
		queryPortfolioById: queryPortfolioById,
		cmdPortfolioCreate: cmdPortfolioCreate,
	}
}

func (h handler) GetPortfolioById(w http.ResponseWriter, r *http.Request, portfolioId PortfolioId) {
	respond := response.New(w, r)

	portfolio, err := h.queryPortfolioById.Fetch(query.PortfolioId(portfolioId))
	if err != nil {
		respond.Err(err)
		return
	}

	respond.Json(http.StatusOK, presenterPortfolio(portfolio))
}

func (h handler) CreatePortfolio(w http.ResponseWriter, r *http.Request) {
	respond := response.New(w, r)

	portfolio := new(PortfolioUpdate)
	if err := respond.Decode(portfolio); err != nil {
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

package portfolio

import (
	"net/http"

	"primedivident/internal/infrastructure/server/http/response"
	"primedivident/internal/modules/portfolio/interactor/query"
	"primedivident/pkg/logger"
)

type handler struct {
	logger             logger.Logger
	queryPortfolioById query.PortfolioById
}

func NewHandler(
	logger logger.Logger,
	queryPortfolioById query.PortfolioById,
) ServerInterface {
	return handler{
		logger:             logger,
		queryPortfolioById: queryPortfolioById,
	}
}

func (h handler) GetPortfolioById(w http.ResponseWriter, r *http.Request, portfolioId PortfolioId) {
	respond := response.New(w, r)

	portfolio, err := h.queryPortfolioById.Fetch(portfolioId)
	if err != nil {
		respond.Err(err)
		return
	}

	respond.Json(http.StatusOK, Portfolio{
		Id:        portfolio.ID,
		CreatedAt: portfolio.CreatedAt,
	})
}

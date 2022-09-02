package portfolio

import (
	"net/http"
	"primedivident/internal/infrastructures/server/http/response"
	"primedivident/pkg/logger"
	"primedivident/pkg/utils"
	"time"
)

type handler struct {
	logger logger.Logger
}

func NewHandler(logger logger.Logger) ServerInterface {
	return handler{logger: logger}
}

func (s handler) GetPortfolioById(w http.ResponseWriter, r *http.Request, portfolioId PortfolioId) {
	response.New(w, r).Json(http.StatusOK, Portfolio{
		CreatedAt: utils.Ptr(time.Now()),
	})
}

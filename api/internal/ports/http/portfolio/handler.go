package portfolio

import (
	"net/http"
	"primedivident/internal/infrastructures/server/http/response"
	"primedivident/pkg/utils"
	"time"
)

type handler struct {
}

func NewHandler() ServerInterface {
	return handler{}
}

func (s handler) GetPortfolioById(w http.ResponseWriter, r *http.Request, portfolioId PortfolioId) {
	response.New(w, r).Json(http.StatusOK, Portfolio{
		CreatedAt: utils.Ptr(time.Now()),
	})
}

package portfolio

import (
	"net/http"
	"primedivident/internal/infrastructures/server/http/response"
	"primedivident/pkg/utils"
	"time"
)

type HttpServer struct {
}

func NewHttpServer() ServerInterface {
	return HttpServer{}
}

func (s HttpServer) GetPortfolioById(w http.ResponseWriter, r *http.Request, portfolioId PortfolioId) {
	response.New(w, r).Json(http.StatusOK, Portfolio{
		CreatedAt: utils.NewOf(time.Now()),
	})
}

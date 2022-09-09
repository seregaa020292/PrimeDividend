package portfolio

import (
	"context"
	"net/http"
	"primedivident/internal/infrastructure/http/openapi"
	"primedivident/pkg/response"
	"primedivident/pkg/validator"
)

var _ openapi.ServerInterface = (*Request)(nil)

type ctxKey int

const (
	portfolioUpdateKey ctxKey = iota
)

type Request struct {
	openapi.ServerInterface
	validator validator.Validator
}

func NewRequest(server openapi.ServerInterface) Request {
	return Request{
		ServerInterface: server,
		validator:       validator.GetValidator(),
	}
}

func (request Request) CreatePortfolio(w http.ResponseWriter, r *http.Request) {
	respond := response.New(w, r)

	portfolio := openapi.PortfolioUpdate{}
	if err := respond.Decode(&portfolio); err != nil {
		respond.Err(err)
		return
	}

	if err := request.validator.Struct(portfolio); err != nil {
		respond.Err(err)
		return
	}

	ctx := context.WithValue(r.Context(), portfolioUpdateKey, portfolio)

	request.ServerInterface.CreatePortfolio(w, r.WithContext(ctx))
}

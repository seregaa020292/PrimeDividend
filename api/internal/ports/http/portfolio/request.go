package portfolio

import (
	"context"
	"net/http"
	"primedivident/internal/infrastructure/http/openapi"
	"primedivident/pkg/response"
	"primedivident/pkg/validator"
)

type ctxKey int

const (
	portfolioUpdateKey ctxKey = iota
)

type Request struct {
	validator validator.Validator
}

func NewRequest(validator validator.Validator) Request {
	return Request{validator: validator}
}

func (request Request) CreatePortfolio(next http.Handler, w http.ResponseWriter, r *http.Request) {
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

	next.ServeHTTP(w, r.WithContext(ctx))
}

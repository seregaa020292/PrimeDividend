package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"primedividend/api/internal/infrastructure/server/openapi"
	"primedividend/api/internal/infrastructure/server/response"
	"primedividend/api/internal/ports/http/asset"
	"primedividend/api/internal/ports/http/auth"
	"primedividend/api/internal/ports/http/currency"
	"primedividend/api/internal/ports/http/instrument"
	"primedividend/api/internal/ports/http/market"
	"primedividend/api/internal/ports/http/portfolio"
	"primedividend/api/internal/ports/http/provider"
	"primedividend/api/internal/ports/http/register"
	"primedividend/api/internal/ports/http/user"
)

// HttpHandlers implements openapi.ServerInterface
var _ openapi.ServerInterface = (*HttpHandlers)(nil) //nolint:typecheck

type HttpHandlers struct {
	auth.HandlerAuth
	asset.HandlerAsset
	currency.HandlerCurrency
	instrument.HandlerInstrument
	market.HandlerMarket
	portfolio.HandlerPortfolio
	provider.HandlerProvider
	register.HandlerRegister
	user.HandlerUser
}

func NewHttpHandlers(
	auth auth.HandlerAuth,
	asset asset.HandlerAsset,
	currency currency.HandlerCurrency,
	instrument instrument.HandlerInstrument,
	market market.HandlerMarket,
	portfolio portfolio.HandlerPortfolio,
	provider provider.HandlerProvider,
	register register.HandlerRegister,
	user user.HandlerUser,
) HttpHandlers {
	return HttpHandlers{
		HandlerAuth:       auth,
		HandlerAsset:      asset,
		HandlerCurrency:   currency,
		HandlerInstrument: instrument,
		HandlerMarket:     market,
		HandlerPortfolio:  portfolio,
		HandlerProvider:   provider,
		HandlerRegister:   register,
		HandlerUser:       user,
	}
}

func (handlers HttpHandlers) Handle(router chi.Router, middlewares []openapi.MiddlewareFunc) {
	openapi.HandlerWithOptions(handlers, openapi.ChiServerOptions{
		BaseRouter:  router,
		Middlewares: middlewares,
		ErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			respond := response.NewRespondBuilder(w, r)
			respond.Err(err)
		},
	})
}

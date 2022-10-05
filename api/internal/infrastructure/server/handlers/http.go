package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"primedivident/internal/infrastructure/server/openapi"
	"primedivident/internal/infrastructure/server/response"
	"primedivident/internal/ports/http/asset"
	"primedivident/internal/ports/http/auth"
	"primedivident/internal/ports/http/currency"
	"primedivident/internal/ports/http/instrument"
	"primedivident/internal/ports/http/market"
	"primedivident/internal/ports/http/portfolio"
	"primedivident/internal/ports/http/provider"
	"primedivident/internal/ports/http/register"
	"primedivident/internal/ports/http/user"
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

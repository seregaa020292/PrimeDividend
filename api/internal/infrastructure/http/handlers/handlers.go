package handlers

import (
	"log"
	"net/http"

	"github.com/getkin/kin-openapi/routers/gorillamux"
	"github.com/go-chi/chi/v5"

	serverHttp "primedivident/internal/infrastructure/http"
	"primedivident/internal/infrastructure/http/openapi"
	"primedivident/internal/modules/auth/service/auth/strategies"
	"primedivident/internal/ports/http/asset"
	"primedivident/internal/ports/http/auth"
	"primedivident/internal/ports/http/currency"
	"primedivident/internal/ports/http/instrument"
	"primedivident/internal/ports/http/market"
	"primedivident/internal/ports/http/portfolio"
	"primedivident/internal/ports/http/provider"
	"primedivident/internal/ports/http/register"
	"primedivident/internal/ports/http/user"
	"primedivident/pkg/response"
)

// Handlers implements openapi.ServerInterface
var _ openapi.ServerInterface = (*Handlers)(nil) //nolint:typecheck

type Handlers struct {
	strategies strategies.Strategies

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

func NewHandlers(
	strategies strategies.Strategies,
	auth auth.HandlerAuth,
	asset asset.HandlerAsset,
	currency currency.HandlerCurrency,
	instrument instrument.HandlerInstrument,
	market market.HandlerMarket,
	portfolio portfolio.HandlerPortfolio,
	provider provider.HandlerProvider,
	register register.HandlerRegister,
	user user.HandlerUser,
) serverHttp.Handlers {
	return Handlers{
		strategies: strategies,

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

func (handlers Handlers) Setup(router chi.Router) {
	swagger, err := openapi.GetSwagger()
	if err != nil {
		log.Fatalln(err)
	}

	swagger.Servers = nil

	router.Use(authValidator(swagger, handlers.strategies))

	routerSwagger, _ := gorillamux.NewRouter(swagger)

	openapi.HandlerWithOptions(handlers, openapi.ChiServerOptions{ //nolint:typecheck
		BaseRouter: router,
		Middlewares: []openapi.MiddlewareFunc{
			custom(routerSwagger),
		},
		ErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			respond := response.NewRespondBuilder(w, r)
			respond.Err(err)
		},
	})
}

package handlers

import (
	"github.com/go-chi/chi/v5"

	serverHttp "primedivident/internal/infrastructure/http"
	"primedivident/internal/infrastructure/http/openapi"
	"primedivident/internal/ports/http/asset"
	"primedivident/internal/ports/http/currency"
	"primedivident/internal/ports/http/instrument"
	"primedivident/internal/ports/http/market"
	"primedivident/internal/ports/http/portfolio"
	"primedivident/internal/ports/http/provider"
	"primedivident/internal/ports/http/register"
	"primedivident/internal/ports/http/user"
)

var _ openapi.ServerInterface = (*Handlers)(nil)

type Handlers struct {
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

func (h Handlers) Setup(router chi.Router) {
	openapi.HandlerFromMux(h, router)
}

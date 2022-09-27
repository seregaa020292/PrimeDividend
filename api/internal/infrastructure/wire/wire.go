//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"

	"primedivident/internal/config"
	serverHttp "primedivident/internal/infrastructure/http"
	"primedivident/internal/infrastructure/http/handlers"
	wireGroup "primedivident/internal/infrastructure/wire/wire_group"
	"primedivident/pkg/response"
	"primedivident/pkg/validator"
)

func Initialize(cfg config.Config) serverHttp.Server {
	wire.Build(
		ProvideLogger,
		ProvidePostgres,
		ProvideMailerObserver,
		ProvideTemplate,
		ProvideJwtTokens,

		validator.GetValidator,
		response.NewRespond,

		wireGroup.Auth,
		wireGroup.Asset,
		wireGroup.Currency,
		wireGroup.Instrument,
		wireGroup.Market,
		wireGroup.Portfolio,
		wireGroup.Provider,
		wireGroup.Register,
		wireGroup.User,

		handlers.NewHandlers,
		serverHttp.NewServer,
	)

	return serverHttp.Server{}
}

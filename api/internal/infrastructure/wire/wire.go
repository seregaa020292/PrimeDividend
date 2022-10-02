//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"

	"primedivident/internal/config"
	"primedivident/internal/infrastructure/response"
	"primedivident/internal/infrastructure/server"
	"primedivident/internal/infrastructure/server/handlers"
	"primedivident/internal/infrastructure/server/middlewares"
	wireGroup "primedivident/internal/infrastructure/wire/wire_group"
	"primedivident/pkg/validator"
)

func Initialize(cfg config.Config) server.Server {
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

		middlewares.NewMiddlewares,
		handlers.NewHandlers,
		server.NewServer,
	)

	return server.Server{}
}

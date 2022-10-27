//go:build wireinject
// +build wireinject

package wire

import (
	"context"

	"github.com/google/wire"

	"primedivident/internal/config"
	"primedivident/internal/handlers"
	"primedivident/internal/infrastructure/server"
	"primedivident/internal/infrastructure/server/response"
	"primedivident/internal/infrastructure/server/routes"
	"primedivident/internal/infrastructure/socket"
	"primedivident/internal/infrastructure/wire/providers"
	wireGroup "primedivident/internal/infrastructure/wire/wire_group"
	"primedivident/pkg/validator"
)

func Initialize(ctx context.Context, cfg config.Config) server.Server {
	wire.Build(
		providers.ProvideLogger,
		providers.ProvidePostgres,
		providers.ProvideRedis,
		providers.ProvideMailerObserver,
		providers.ProvideTemplate,
		providers.ProvideJwtTokens,
		providers.ProvideShutdown,

		socket.NewUpgrader,
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

		handlers.NewWsHandlers,
		handlers.NewHttpHandlers,
		routes.NewRoutes,
		server.NewServer,
	)

	return server.Server{}
}

//go:build wireinject
// +build wireinject

package wire

import (
	"context"

	"github.com/google/wire"

	"primedividend/api/internal/config"
	"primedividend/api/internal/handlers"
	"primedividend/api/internal/infrastructure/server"
	"primedividend/api/internal/infrastructure/server/response"
	"primedividend/api/internal/infrastructure/server/routes"
	"primedividend/api/internal/infrastructure/socket"
	"primedividend/api/internal/infrastructure/wire/providers"
	wireGroup "primedividend/api/internal/infrastructure/wire/wire_group"
	"primedividend/api/pkg/validator"
)

func Initialize(ctx context.Context, cfg config.Config) server.Server {
	wire.Build(
		providers.ProvideLogger,
		providers.ProvidePostgres,
		providers.ProvideTransactor,
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

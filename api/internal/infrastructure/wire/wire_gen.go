// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"primedivident/internal/config"
	"primedivident/internal/infrastructure/response"
	"primedivident/internal/infrastructure/server"
	"primedivident/internal/infrastructure/server/handlers"
	"primedivident/internal/infrastructure/server/middlewares"
	"primedivident/internal/infrastructure/wire/wire_group"
	"primedivident/internal/modules/auth/command"
	repository2 "primedivident/internal/modules/auth/repository"
	"primedivident/internal/modules/auth/service/email"
	"primedivident/internal/modules/auth/service/strategy"
	"primedivident/internal/modules/auth/service/strategy/repository"
	"primedivident/internal/modules/instrument/query"
	repository3 "primedivident/internal/modules/instrument/repository"
	command2 "primedivident/internal/modules/portfolio/command"
	query2 "primedivident/internal/modules/portfolio/query"
	repository4 "primedivident/internal/modules/portfolio/repository"
	"primedivident/internal/ports/http/asset"
	"primedivident/internal/ports/http/auth"
	"primedivident/internal/ports/http/currency"
	instrument2 "primedivident/internal/ports/http/instrument"
	"primedivident/internal/ports/http/market"
	portfolio2 "primedivident/internal/ports/http/portfolio"
	"primedivident/internal/ports/http/provider"
	"primedivident/internal/ports/http/register"
	"primedivident/internal/ports/http/user"
	"primedivident/internal/presenters/instrument"
	"primedivident/internal/presenters/portfolio"
	"primedivident/pkg/validator"
)

// Injectors from wire.go:

func Initialize(cfg config.Config) server.Server {
	serverMiddlewares := middlewares.NewMiddlewares(cfg)
	jwtTokens := ProvideJwtTokens(cfg)
	postgres := ProvidePostgres(cfg)
	repositoryRepository := repository.NewRepository(postgres)
	service := strategy.NewService(jwtTokens, repositoryRepository)
	strategyStrategy := wire_group.ProvideStrategy(cfg, service)
	logger := ProvideLogger(cfg)
	validatorValidator := validator.GetValidator()
	responder := response.NewRespond(logger, validatorValidator)
	repository5 := repository2.NewRepository(postgres)
	sender := ProvideMailerObserver(cfg, logger)
	templater := ProvideTemplate(cfg)
	joinConfirmUser := email.NewJoinConfirmUser(sender, templater)
	joinByEmail := command.NewJoinByEmail(repository5, joinConfirmUser)
	confirmUser := email.NewConfirmUser(sender, templater)
	confirmByToken := command.NewConfirmByToken(repository5, confirmUser)
	handlerAuth := auth.NewHandler(responder, strategyStrategy, joinByEmail, confirmByToken)
	handlerAsset := asset.NewHandler()
	handlerCurrency := currency.NewHandler()
	presenter := instrument.NewPresenter()
	repository6 := repository3.NewRepository(postgres)
	instrumentAll := query.NewInstrumentAll(repository6)
	handlerInstrument := instrument2.NewHandler(responder, presenter, instrumentAll)
	handlerMarket := market.NewHandler()
	portfolioPresenter := portfolio.NewPresenter()
	repository7 := repository4.NewRepository(postgres)
	portfolioById := query2.NewPortfolioById(repository7)
	portfolioCreate := command2.NewPortfolioCreate(repository7)
	handlerPortfolio := portfolio2.NewHandler(responder, portfolioPresenter, portfolioById, portfolioCreate)
	handlerProvider := provider.NewHandler()
	handlerRegister := register.NewHandler()
	handlerUser := user.NewHandler()
	serverHandlers := handlers.NewHandlers(strategyStrategy, handlerAuth, handlerAsset, handlerCurrency, handlerInstrument, handlerMarket, handlerPortfolio, handlerProvider, handlerRegister, handlerUser)
	serverServer := server.NewServer(serverMiddlewares, serverHandlers)
	return serverServer
}

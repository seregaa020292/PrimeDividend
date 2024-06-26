// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"context"
	"primedividend/api/internal/config"
	"primedividend/api/internal/handlers"
	"primedividend/api/internal/infrastructure/server"
	"primedividend/api/internal/infrastructure/server/response"
	"primedividend/api/internal/infrastructure/server/routes"
	"primedividend/api/internal/infrastructure/socket"
	"primedividend/api/internal/infrastructure/wire/providers"
	command2 "primedividend/api/internal/modules/asset/command"
	"primedividend/api/internal/modules/asset/query"
	repository4 "primedividend/api/internal/modules/asset/repository"
	"primedividend/api/internal/modules/auth/command"
	repository3 "primedividend/api/internal/modules/auth/repository"
	"primedividend/api/internal/modules/auth/service/email"
	"primedividend/api/internal/modules/auth/service/strategy"
	repository2 "primedividend/api/internal/modules/auth/service/strategy/repository"
	query2 "primedividend/api/internal/modules/currency/query"
	repository5 "primedividend/api/internal/modules/currency/repository"
	query3 "primedividend/api/internal/modules/instrument/query"
	repository6 "primedividend/api/internal/modules/instrument/repository"
	query4 "primedividend/api/internal/modules/market/query"
	"primedividend/api/internal/modules/market/repository"
	"primedividend/api/internal/modules/market/service/quotes"
	command3 "primedividend/api/internal/modules/portfolio/command"
	query5 "primedividend/api/internal/modules/portfolio/query"
	repository7 "primedividend/api/internal/modules/portfolio/repository"
	query6 "primedividend/api/internal/modules/provider/query"
	repository8 "primedividend/api/internal/modules/provider/repository"
	command4 "primedividend/api/internal/modules/user/command"
	query7 "primedividend/api/internal/modules/user/query"
	repository9 "primedividend/api/internal/modules/user/repository"
	asset2 "primedividend/api/internal/ports/http/asset"
	"primedividend/api/internal/ports/http/auth"
	currency2 "primedividend/api/internal/ports/http/currency"
	instrument2 "primedividend/api/internal/ports/http/instrument"
	market2 "primedividend/api/internal/ports/http/market"
	portfolio2 "primedividend/api/internal/ports/http/portfolio"
	provider2 "primedividend/api/internal/ports/http/provider"
	"primedividend/api/internal/ports/http/register"
	user2 "primedividend/api/internal/ports/http/user"
	market3 "primedividend/api/internal/ports/ws/market"
	"primedividend/api/internal/presenters/asset"
	"primedividend/api/internal/presenters/currency"
	"primedividend/api/internal/presenters/instrument"
	"primedividend/api/internal/presenters/market"
	"primedividend/api/internal/presenters/portfolio"
	"primedividend/api/internal/presenters/provider"
	"primedividend/api/internal/presenters/user"
	"primedividend/api/pkg/validator"
)

// Injectors from wire.go:

func Initialize(ctx context.Context, cfg config.Config) server.Server {
	postgres := providers.ProvidePostgres(cfg)
	redis := providers.ProvideRedis(cfg)
	logger := providers.ProvideLogger(cfg)
	sender := providers.ProvideMailerObserver(cfg, logger)
	assetRepository := repository.NewAssetRepository(redis)
	tinkoff := providers.ProvideTinkoff(cfg)
	hubQuotes := quotes.NewHubQuotes(assetRepository, tinkoff)
	shutdownApp := providers.ProvideShutdown(postgres, redis, sender, hubQuotes)
	jwtTokens := providers.ProvideJwtTokens(cfg)
	repositoryRepository := repository2.NewRepository(postgres)
	service := strategy.NewService(jwtTokens, repositoryRepository)
	strategyStrategy := providers.ProvideStrategy(cfg, service)
	validatorValidator := validator.GetValidator()
	responder := response.NewRespond(logger, validatorValidator)
	repository10 := repository3.NewRepository(postgres)
	templater := providers.ProvideTemplate(cfg)
	joinConfirmUser := email.NewJoinConfirmUser(sender, templater)
	joinByEmail := command.NewJoinByEmail(repository10, joinConfirmUser)
	confirmUser := email.NewConfirmUser(sender, templater)
	confirmByToken := command.NewConfirmByToken(repository10, confirmUser)
	handlerAuth := auth.NewHandler(responder, strategyStrategy, joinByEmail, confirmByToken)
	presenter := asset.NewPresenter()
	repository11 := repository4.NewRepository(postgres)
	getUserAll := query.NewGetUserAll(repository11)
	create := command2.NewCreate(repository11)
	txManager := providers.ProvideTransactor(postgres)
	edit := command2.NewEdit(repository11, txManager)
	remove := command2.NewRemove(repository11)
	handlerAsset := asset2.NewHandler(responder, presenter, getUserAll, create, edit, remove)
	currencyPresenter := currency.NewPresenter()
	repository12 := repository5.NewRepository(postgres)
	getById := query2.NewGetById(repository12)
	getAll := query2.NewGetAll(repository12)
	handlerCurrency := currency2.NewHandler(responder, currencyPresenter, getById, getAll)
	instrumentPresenter := instrument.NewPresenter()
	repository13 := repository6.NewRepository(postgres)
	queryGetById := query3.NewGetById(repository13)
	queryGetAll := query3.NewGetAll(repository13)
	handlerInstrument := instrument2.NewHandler(responder, instrumentPresenter, queryGetById, queryGetAll)
	marketPresenter := market.NewPresenter()
	repository14 := repository.NewRepository(postgres)
	getById2 := query4.NewGetById(repository14)
	getByTicker := query4.NewGetByTicker(repository14)
	getAll2 := query4.NewGetAll(repository14)
	handlerMarket := market2.NewHandler(responder, marketPresenter, getById2, getByTicker, getAll2)
	portfolioPresenter := portfolio.NewPresenter()
	repository15 := repository7.NewRepository(postgres)
	getById3 := query5.NewGetById(repository15)
	getAll3 := query5.NewGetAll(repository15)
	queryGetUserAll := query5.NewGetUserAll(repository15)
	commandCreate := command3.NewCreate(repository15)
	commandEdit := command3.NewEdit(repository15)
	commandRemove := command3.NewRemove(repository15)
	handlerPortfolio := portfolio2.NewHandler(responder, portfolioPresenter, getById3, getAll3, queryGetUserAll, commandCreate, commandEdit, commandRemove)
	providerPresenter := provider.NewPresenter()
	repository16 := repository8.NewRepository(postgres)
	getById4 := query6.NewGetById(repository16)
	getAll4 := query6.NewGetAll(repository16)
	handlerProvider := provider2.NewHandler(responder, providerPresenter, getById4, getAll4)
	handlerRegister := register.NewHandler()
	userPresenter := user.NewPresenter()
	repository17 := repository9.NewRepository(postgres)
	getById5 := query7.NewGetById(repository17)
	remove2 := command4.NewRemove(repository17)
	edit2 := command4.NewEdit(repository17)
	handlerUser := user2.NewHandler(responder, userPresenter, getById5, remove2, edit2)
	httpHandlers := handlers.NewHttpHandlers(handlerAuth, handlerAsset, handlerCurrency, handlerInstrument, handlerMarket, handlerPortfolio, handlerProvider, handlerRegister, handlerUser)
	upgrader := socket.NewUpgrader()
	marketHandlerMarket := market3.NewHandlerMarket(responder, upgrader, hubQuotes)
	wsHandlers := handlers.NewWsHandlers(marketHandlerMarket)
	routesRoutes := routes.NewRoutes(strategyStrategy, httpHandlers, wsHandlers)
	serverServer := server.NewServer(ctx, shutdownApp, routesRoutes)
	return serverServer
}

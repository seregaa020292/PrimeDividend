// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"primedivident/internal/config"
	"primedivident/internal/infrastructure/http"
	"primedivident/internal/infrastructure/http/handlers"
	"primedivident/internal/modules/auth/command"
	"primedivident/internal/modules/auth/repository"
	"primedivident/internal/modules/auth/service/email"
	"primedivident/internal/modules/instrument/query"
	repository2 "primedivident/internal/modules/instrument/repository"
	command2 "primedivident/internal/modules/portfolio/command"
	query2 "primedivident/internal/modules/portfolio/query"
	repository3 "primedivident/internal/modules/portfolio/repository"
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
	"primedivident/pkg/validator"
)

// Injectors from wire.go:

func Initialize(cfg config.Config) http.Server {
	logger := ProvideLogger(cfg)
	validatorValidator := validator.GetValidator()
	responder := response.NewRespond(logger, validatorValidator)
	postgres := ProvidePostgres(cfg)
	repositoryRepository := repository.NewRepository(postgres)
	sender := ProvideMailerObserver(cfg, logger)
	joinConfirmUser := email.NewJoinConfirmUser(sender)
	joinByEmail := command.NewJoinByEmail(repositoryRepository, joinConfirmUser)
	confirmByToken := command.NewConfirmByToken(repositoryRepository, joinConfirmUser)
	handlerAuth := auth.NewHandler(responder, joinByEmail, confirmByToken)
	handlerAsset := asset.NewHandler()
	handlerCurrency := currency.NewHandler()
	repository4 := repository2.NewRepository(postgres)
	instrumentAll := query.NewInstrumentAll(repository4)
	handlerInstrument := instrument.NewHandler(responder, instrumentAll)
	handlerMarket := market.NewHandler()
	repository5 := repository3.NewRepository(postgres)
	portfolioById := query2.NewPortfolioById(repository5)
	portfolioCreate := command2.NewPortfolioCreate(repository5)
	handlerPortfolio := portfolio.NewHandler(responder, portfolioById, portfolioCreate)
	handlerProvider := provider.NewHandler()
	handlerRegister := register.NewHandler()
	handlerUser := user.NewHandler()
	httpHandlers := handlers.NewHandlers(handlerAuth, handlerAsset, handlerCurrency, handlerInstrument, handlerMarket, handlerPortfolio, handlerProvider, handlerRegister, handlerUser)
	server := http.NewServer(httpHandlers)
	return server
}

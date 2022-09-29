package wire_group

import (
	"github.com/google/wire"

	"primedivident/internal/config"
	"primedivident/internal/modules/auth/command"
	"primedivident/internal/modules/auth/repository"
	"primedivident/internal/modules/auth/service/email"
	"primedivident/internal/modules/auth/service/strategy"
	"primedivident/internal/modules/auth/service/strategy/auth"
	repository2 "primedivident/internal/modules/auth/service/strategy/repository"
	"primedivident/internal/modules/auth/service/strategy/strategies"
	port "primedivident/internal/ports/http/auth"
)

func ProvideStrategy(cfg config.Config, jwtTokens auth.JwtTokens, repository repository2.Repository) strategy.Strategy {
	newStrategy := strategy.NewStrategy(jwtTokens, repository)
	strategyService := strategy.NewService(jwtTokens, repository)

	newStrategy.Password().Set(auth.Email, strategies.NewEmailStrategy(strategyService))
	newStrategy.Network().Set(auth.Vk, strategies.NewVkStrategy(cfg.Networks.VkOAuth2, strategyService))
	newStrategy.Network().Set(auth.Ok, strategies.NewOkStrategy(cfg.Networks.OkOAuth2, strategyService))
	newStrategy.Network().Set(auth.Yandex, strategies.NewYandexStrategy(cfg.Networks.YandexOAuth2, strategyService))

	return newStrategy
}

var Auth = wire.NewSet(
	repository.NewRepository,
	repository2.NewRepository,
	ProvideStrategy,
	email.NewJoinConfirmUser,
	email.NewConfirmUser,
	command.NewJoinByEmail,
	command.NewConfirmByToken,
	port.NewHandler,
)

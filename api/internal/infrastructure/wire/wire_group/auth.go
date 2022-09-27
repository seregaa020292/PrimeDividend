package wire_group

import (
	"github.com/google/wire"

	"primedivident/internal/config"
	"primedivident/internal/modules/auth/command"
	"primedivident/internal/modules/auth/repository"
	"primedivident/internal/modules/auth/service/email"
	"primedivident/internal/modules/auth/service/strategy"
	"primedivident/internal/modules/auth/service/strategy/auth"
	sr "primedivident/internal/modules/auth/service/strategy/repository"
	"primedivident/internal/modules/auth/service/strategy/strategies"
	port "primedivident/internal/ports/http/auth"
)

func ProvideStrategy(cfg config.Config, jwtTokens auth.JwtTokens, repository sr.Repository) strategy.Strategy {
	s := strategy.NewStrategy(jwtTokens, repository)

	s.Password().Set(auth.Email, strategies.NewEmailStrategy(jwtTokens, repository))
	s.Network().Set(auth.Vk, strategies.NewVkStrategy(cfg.Networks.VkOAuth2, jwtTokens, repository))
	s.Network().Set(auth.Ok, strategies.NewOkStrategy(cfg.Networks.OkOAuth2, jwtTokens, repository))
	s.Network().Set(auth.Yandex, strategies.NewYandexStrategy(cfg.Networks.YandexOAuth2, jwtTokens, repository))

	return s
}

var Auth = wire.NewSet(
	repository.NewRepository,
	sr.NewRepository,
	ProvideStrategy,
	email.NewJoinConfirmUser,
	email.NewConfirmUser,
	command.NewJoinByEmail,
	command.NewConfirmByToken,
	port.NewHandler,
)

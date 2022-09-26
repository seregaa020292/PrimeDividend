package wire_group

import (
	"github.com/google/wire"

	"primedivident/internal/config"
	"primedivident/internal/modules/auth/command"
	"primedivident/internal/modules/auth/repository"
	"primedivident/internal/modules/auth/service/auth"
	"primedivident/internal/modules/auth/service/auth/strategies"
	"primedivident/internal/modules/auth/service/email"
	port "primedivident/internal/ports/http/auth"
)

func ProvideAuth(cfg config.Config, jwtTokens auth.JwtTokens, repository repository.Repository) auth.Auth {
	strategy := auth.NewStrategy()

	strategy.SetPassword(auth.Email, strategies.NewEmailStrategy(jwtTokens, repository))
	strategy.SetNetwork(auth.Vk, strategies.NewVkStrategy(cfg.Networks.VkOAuth2, jwtTokens, repository))
	strategy.SetNetwork(auth.Ok, strategies.NewOkStrategy(cfg.Networks.OkOAuth2, jwtTokens, repository))
	strategy.SetNetwork(auth.Yandex, strategies.NewYandexStrategy(cfg.Networks.YandexOAuth2, jwtTokens, repository))

	return auth.NewAuth(strategy, jwtTokens, repository)
}

var Auth = wire.NewSet(
	repository.NewRepository,
	email.NewJoinConfirmUser,
	email.NewConfirmUser,
	command.NewJoinByEmail,
	command.NewConfirmByToken,
	port.NewHandler,
)

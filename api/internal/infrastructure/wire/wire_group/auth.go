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

func ProvideStrategies(cfg config.Config, repository repository.Repository) strategies.Strategies {
	jwtTokens := auth.NewJwtTokens(cfg.App.Name, cfg.Jwt)
	return strategies.NewStrategies(cfg.Networks, jwtTokens, repository)
}

var Auth = wire.NewSet(
	repository.NewRepository,
	ProvideStrategies,
	email.NewJoinConfirmUser,
	email.NewConfirmUser,
	command.NewJoinByEmail,
	command.NewConfirmByToken,
	port.NewHandler,
)

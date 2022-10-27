package wire_group

import (
	"github.com/google/wire"

	"primedivident/internal/infrastructure/wire/providers"
	"primedivident/internal/modules/auth/command"
	"primedivident/internal/modules/auth/repository"
	"primedivident/internal/modules/auth/service/email"
	"primedivident/internal/modules/auth/service/strategy"
	strategyRepository "primedivident/internal/modules/auth/service/strategy/repository"
	http "primedivident/internal/ports/http/auth"
)

var Auth = wire.NewSet(
	repository.NewRepository,
	strategyRepository.NewRepository,
	strategy.NewService,
	providers.ProvideStrategy,
	email.NewJoinConfirmUser,
	email.NewConfirmUser,
	command.NewJoinByEmail,
	command.NewConfirmByToken,
	http.NewHandler,
)

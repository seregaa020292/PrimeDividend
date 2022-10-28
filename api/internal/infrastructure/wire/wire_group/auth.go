package wire_group

import (
	"github.com/google/wire"

	"primedividend/api/internal/infrastructure/wire/providers"
	"primedividend/api/internal/modules/auth/command"
	"primedividend/api/internal/modules/auth/repository"
	"primedividend/api/internal/modules/auth/service/email"
	"primedividend/api/internal/modules/auth/service/strategy"
	strategyRepository "primedividend/api/internal/modules/auth/service/strategy/repository"
	http "primedividend/api/internal/ports/http/auth"
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

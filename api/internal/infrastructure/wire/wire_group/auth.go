package wire_group

import (
	"github.com/google/wire"

	"primedivident/internal/modules/auth/interactor/command"
	"primedivident/internal/modules/auth/repository"
	"primedivident/internal/modules/auth/service/email"
	"primedivident/internal/ports/http/auth"
)

var Auth = wire.NewSet(
	email.NewJoinConfirmUser,
	repository.NewRepository,
	command.NewJoinByEmail,
	command.NewConfirmByToken,
	auth.NewHandler,
)

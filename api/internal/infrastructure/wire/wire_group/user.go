package wire_group

import (
	"github.com/google/wire"

	"primedivident/internal/modules/user/command"
	"primedivident/internal/modules/user/query"
	"primedivident/internal/modules/user/repository"
	"primedivident/internal/ports/http/user"
	presenter "primedivident/internal/presenters/user"
)

var User = wire.NewSet(
	presenter.NewPresenter,
	repository.NewRepository,
	query.NewGetById,
	command.NewRemove,
	command.NewEdit,
	user.NewHandler,
)

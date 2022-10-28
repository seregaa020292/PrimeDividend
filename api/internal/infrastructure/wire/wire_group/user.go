package wire_group

import (
	"github.com/google/wire"

	"primedividend/api/internal/modules/user/command"
	"primedividend/api/internal/modules/user/query"
	"primedividend/api/internal/modules/user/repository"
	"primedividend/api/internal/ports/http/user"
	presenter "primedividend/api/internal/presenters/user"
)

var User = wire.NewSet(
	presenter.NewPresenter,
	repository.NewRepository,
	query.NewGetById,
	command.NewRemove,
	command.NewEdit,
	user.NewHandler,
)

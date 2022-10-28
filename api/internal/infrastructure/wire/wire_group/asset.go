package wire_group

import (
	"github.com/google/wire"

	"primedividend/api/internal/modules/asset/command"
	"primedividend/api/internal/modules/asset/query"
	"primedividend/api/internal/modules/asset/repository"
	"primedividend/api/internal/ports/http/asset"
	presenter "primedividend/api/internal/presenters/asset"
)

var Asset = wire.NewSet(
	presenter.NewPresenter,
	repository.NewRepository,
	query.NewGetUserAll,
	command.NewCreate,
	command.NewEdit,
	command.NewRemove,
	asset.NewHandler,
)

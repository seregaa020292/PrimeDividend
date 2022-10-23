package wire_group

import (
	"github.com/google/wire"

	"primedivident/internal/modules/asset/query"
	"primedivident/internal/modules/asset/repository"
	"primedivident/internal/ports/http/asset"
	presenter "primedivident/internal/presenters/asset"
)

var Asset = wire.NewSet(
	presenter.NewPresenter,
	repository.NewRepository,
	query.NewGetUserAll,
	asset.NewHandler,
)

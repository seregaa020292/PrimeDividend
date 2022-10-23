package asset

import (
	"net/http"

	"primedivident/internal/infrastructure/server/middlewares/helper"
	"primedivident/internal/infrastructure/server/openapi"
	"primedivident/internal/modules/asset/query"
)

func (h HandlerAsset) GetUserAssets(w http.ResponseWriter, r *http.Request, params openapi.GetUserAssetsParams) {
	respond := h.responder.Http(w, r)

	portfolioID, err := helper.UUID(params.PortfolioId)
	if err != nil {
		respond.Err(err)
		return
	}

	user, err := helper.UserFromCtx(r.Context())
	if err != nil {
		respond.Err(err)
		return
	}

	assets, err := h.queryGetUserAll.Fetch(query.PayloadUserAll{
		UserID:      user.ID,
		PortfolioID: portfolioID,
	})
	if err != nil {
		respond.Err(err)
		return
	}

	respond.Json(http.StatusOK, h.presenter.GetAll(assets))
}

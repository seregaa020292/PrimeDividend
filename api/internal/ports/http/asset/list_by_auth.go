package asset

import (
	"net/http"

	"primedivident/internal/infrastructure/server/middlewares/helper"
	"primedivident/internal/infrastructure/server/openapi"
	"primedivident/internal/modules/asset/query"
)

func (h HandlerAsset) GetUserAssets(w http.ResponseWriter, r *http.Request, portfolioId openapi.PortfolioId) {
	respond := h.responder.Http(w, r)

	user, err := helper.UserFromCtx(r.Context())
	if err != nil {
		respond.Err(err)
		return
	}

	assets, err := h.queryGetUserAll.Fetch(query.PayloadUserAll{
		UserID:      user.ID,
		PortfolioID: portfolioId,
	})
	if err != nil {
		respond.Err(err)
		return
	}

	respond.Json(http.StatusOK, h.presenter.GetAll(assets))
}

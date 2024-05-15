package asset

import (
	"net/http"

	"primedividend/api/internal/infrastructure/server/middlewares/helper"
	"primedividend/api/internal/infrastructure/server/openapi"
	"primedividend/api/internal/modules/asset/query"
)

func (h HandlerAsset) GetUserAssets(w http.ResponseWriter, r *http.Request, portfolioId openapi.PortfolioId) {
	respond := h.responder.Http(w, r)

	user, err := helper.UserFromCtx(r.Context())
	if err != nil {
		respond.Err(err)
		return
	}

	assets, err := h.queryGetUserAll.Fetch(r.Context(), query.PayloadUserAll{
		UserID:      user.ID,
		PortfolioID: portfolioId,
	})
	if err != nil {
		respond.Err(err)
		return
	}

	respond.Json(http.StatusOK, h.presenter.GetAll(assets))
}

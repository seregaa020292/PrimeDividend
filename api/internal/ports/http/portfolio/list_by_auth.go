package portfolio

import (
	"net/http"

	"primedividend/api/internal/infrastructure/server/middlewares/helper"
	"primedividend/api/internal/modules/portfolio/query"
	"primedividend/api/pkg/utils/gog"
)

func (h HandlerPortfolio) GetUserPortfolios(w http.ResponseWriter, r *http.Request) {
	respond := h.responder.Http(w, r)

	user, err := helper.UserFromCtx(r.Context())
	if err != nil {
		respond.Err(err)
		return
	}

	portfolios, err := h.queryGetUserAll.Fetch(query.PayloadUserAll{
		UserID: user.ID,
		Active: gog.Ptr(true),
	})
	if err != nil {
		respond.Err(err)
		return
	}

	respond.Json(http.StatusOK, h.presenter.GetAll(portfolios))
}

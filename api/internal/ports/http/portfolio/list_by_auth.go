package portfolio

import (
	"net/http"

	"primedivident/internal/infrastructure/server/middlewares/helper"
	"primedivident/internal/modules/portfolio/query"
	"primedivident/pkg/utils/gog"
)

func (h HandlerPortfolio) GetUserPortfolios(w http.ResponseWriter, r *http.Request) {
	respond := h.responder.Http(w, r)

	user, err := helper.UserFromCtx(r.Context())
	if err != nil {
		respond.Err(err)
		return
	}

	portfolios, err := h.queryGetUserAll.Fetch(query.FilterGetUserAll{
		UserID: user.ID,
		Active: gog.Ptr(true),
	})
	if err != nil {
		respond.Err(err)
		return
	}

	respond.Json(http.StatusOK, h.presenter.GetAll(portfolios))
}

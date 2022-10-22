package user

import (
	"net/http"

	"primedivident/internal/infrastructure/server/openapi"
)

func (h HandlerUser) GetUserById(w http.ResponseWriter, r *http.Request, userId openapi.UserId) {
	respond := h.responder.Http(w, r)

	portfolio, err := h.queryGetById.Fetch(userId)
	if err != nil {
		respond.Err(err)
		return
	}

	respond.Json(http.StatusOK, h.presenter.GetOne(portfolio))
}

package user

import (
	"net/http"

	"primedivident/internal/infrastructure/server/middlewares/helper"
)

func (h HandlerUser) GetUser(w http.ResponseWriter, r *http.Request) {
	respond := h.responder.Http(w, r)

	auth, err := helper.UserFromCtx(r.Context())
	if err != nil {
		respond.Err(err)
		return
	}

	user, err := h.queryGetById.Fetch(auth.ID)
	if err != nil {
		respond.Err(err)
		return
	}

	respond.Json(http.StatusOK, h.presenter.GetOne(user))
}

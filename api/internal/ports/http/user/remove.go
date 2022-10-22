package user

import (
	"net/http"

	"primedivident/internal/infrastructure/server/middlewares/helper"
)

func (h HandlerUser) RemoveUser(w http.ResponseWriter, r *http.Request) {
	respond := h.responder.Http(w, r)

	auth, err := helper.UserFromCtx(r.Context())
	if err != nil {
		respond.Err(err)
		return
	}

	if err := h.commandRemove.Exec(auth.ID); err != nil {
		respond.Err(err)
		return
	}

	respond.WriteHeader(http.StatusOK)
}

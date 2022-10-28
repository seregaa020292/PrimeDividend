package user

import (
	"net/http"

	"primedividend/api/internal/infrastructure/server/middlewares/helper"
	"primedividend/api/internal/infrastructure/server/openapi"
	"primedividend/api/internal/modules/user/command"
)

func (h HandlerUser) UpdateUser(w http.ResponseWriter, r *http.Request) {
	respond := h.responder.Http(w, r)

	auth, err := helper.UserFromCtx(r.Context())
	if err != nil {
		respond.Err(err)
		return
	}

	var user openapi.UserUpdate

	if err := respond.DecodeValidate(&user); err != nil {
		respond.Err(err)
		return
	}

	if err := h.commandEdit.Exec(command.PayloadEdit{
		UserID: auth.ID,
		Name:   user.Name,
		Email:  user.Email,
	}); err != nil {
		respond.Err(err)
		return
	}

	respond.WriteHeader(http.StatusOK)
}

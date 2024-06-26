package auth

import (
	"net/http"

	"primedividend/api/internal/infrastructure/server/openapi"
	"primedividend/api/internal/modules/auth/command"
)

func (h HandlerAuth) JoinEmail(w http.ResponseWriter, r *http.Request) {
	respond := h.responder.Http(w, r)

	var user openapi.AuthUser

	if err := respond.DecodeValidate(&user); err != nil {
		respond.Err(err)
		return
	}

	if err := h.cmdJoinByEmail.Exec(command.Credential(user)); err != nil {
		respond.Err(err)
		return
	}

	respond.WriteHeader(http.StatusCreated)
}

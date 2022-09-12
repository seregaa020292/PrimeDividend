package auth

import (
	"net/http"

	"primedivident/internal/infrastructure/http/openapi"
	"primedivident/internal/modules/auth/interactor/command"
)

func (h HandlerAuth) AuthEmail(w http.ResponseWriter, r *http.Request) {
	respond := h.responder.Http(w, r)

	user := openapi.AuthUser{}

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

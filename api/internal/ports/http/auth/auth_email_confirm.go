package auth

import (
	"net/http"

	"primedivident/internal/infrastructure/http/openapi"
)

func (h HandlerAuth) AuthEmailConfirm(w http.ResponseWriter, r *http.Request) {
	respond := h.responder.Http(w, r)

	confirm := openapi.AuthConfirm{}

	if err := respond.DecodeValidate(&confirm); err != nil {
		respond.Err(err)
		return
	}

	if err := h.cmdConfirmByToken.Exec(confirm.Token); err != nil {
		respond.Err(err)
		return
	}

	respond.WriteHeader(http.StatusOK)
}

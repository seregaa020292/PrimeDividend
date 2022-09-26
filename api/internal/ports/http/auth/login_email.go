package auth

import (
	"net/http"

	"primedivident/internal/infrastructure/http/openapi"
	"primedivident/internal/modules/auth/service/auth"
)

func (h HandlerAuth) LoginEmail(w http.ResponseWriter, r *http.Request) {
	respond := h.responder.Http(w, r)

	var user openapi.LoginUser
	if err := respond.DecodeValidate(&user); err != nil {
		respond.Err(err)
		return
	}

	tokens, err := h.authService.PasswordStrategy(auth.Email).Login(user.Email, user.Password)
	if err != nil {
		respond.Err(err)
		return
	}

	respond.Json(http.StatusOK, tokens)
}

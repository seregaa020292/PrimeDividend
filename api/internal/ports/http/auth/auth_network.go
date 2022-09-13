package auth

import (
	"net/http"

	"golang.org/x/oauth2"

	"primedivident/internal/infrastructure/auth"
	"primedivident/internal/infrastructure/http/openapi"
)

func (h HandlerAuth) AuthNetwork(w http.ResponseWriter, r *http.Request, network openapi.Network) {
	respond := h.responder.Http(w, r)

	url := auth.VkOAuth2Config.AuthCodeURL(auth.GenStateOauthCookie(w, r), oauth2.AccessTypeOnline)

	respond.Redirect(url)
}

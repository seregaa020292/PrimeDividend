package auth

import (
	"errors"
	"net/http"

	"github.com/google/uuid"

	"primedivident/internal/config/consts"
	"primedivident/pkg/datetime"
	"primedivident/pkg/utils"
)

func ValidateOauthState(r *http.Request) error {
	oauthState, err := r.Cookie(consts.OauthState)
	if err != nil {
		return err
	}

	if r.FormValue("state") != oauthState.Value {
		return errors.New("invalid oauth state")
	}

	return nil
}

func GenCookieOauthState(w http.ResponseWriter, r *http.Request) string {
	state := uuid.New().String()

	cookie := utils.GenCookie(consts.OauthState, state, &http.Cookie{
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Domain:   r.URL.Hostname(),
		Expires:  datetime.GetNow().AddDate(1, 0, 0),
	})

	http.SetCookie(w, cookie)

	return state
}

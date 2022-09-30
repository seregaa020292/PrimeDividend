package strategy

import (
	"errors"
	"net/http"

	"github.com/google/uuid"

	"primedivident/pkg/datetime"
	"primedivident/pkg/utils"
)

const OauthState = "oauth-state"

func ValidateOauthState(state string, r *http.Request) error {
	oauthState, err := r.Cookie(OauthState)
	if err != nil {
		return err
	}

	if state != oauthState.Value {
		return errors.New("invalid oauth state")
	}

	return nil
}

func GenCookieOauthState(w http.ResponseWriter, r *http.Request) string {
	state := uuid.New().String()

	cookie := utils.GenCookie(OauthState, state, &http.Cookie{
		Secure:   true,
		HttpOnly: true,
		Domain:   r.URL.Hostname(),
		Expires:  datetime.GetNow().AddDate(1, 0, 0),
	})

	http.SetCookie(w, cookie)

	return state
}

package strategy

import (
	"errors"
	"net/http"

	"github.com/google/uuid"

	"primedivident/pkg/datetime"
	"primedivident/pkg/errs"
	"primedivident/pkg/errs/errmsg"
	"primedivident/pkg/utils"
)

const OauthState = "oauth-state"

func ValidateOauthState(state string, r *http.Request) error {
	oauthState, err := r.Cookie(OauthState)
	if err != nil {
		return errs.BadRequest.Wrap(err, errmsg.FailedGetData)
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

func RemoveCookieOauthState(w http.ResponseWriter, r *http.Request) {
	cookie := utils.GenCookie(OauthState, "", &http.Cookie{
		Domain: r.URL.Hostname(),
		MaxAge: -1,
	})

	http.SetCookie(w, cookie)
}

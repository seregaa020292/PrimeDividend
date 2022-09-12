package middlewares

import (
	"net/http"

	"primedivident/pkg/response"
)

func Recovered(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil && rvr != http.ErrAbortHandler {
				respond := response.NewRespondBuilder(w, r)
				respond.Err(rvr.(error))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

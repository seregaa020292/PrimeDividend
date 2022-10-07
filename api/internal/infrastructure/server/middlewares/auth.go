package middlewares

import (
	"net/http"

	"primedivident/internal/infrastructure/server/middlewares/helper"
	"primedivident/internal/infrastructure/server/response"
	"primedivident/internal/modules/auth/service/strategy"
)

type Auth struct {
	Strategy strategy.Strategy
}

func (a Auth) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		respond := response.NewRespondBuilder(w, r)

		token, err := helper.TokenFromRequest(r)
		if err != nil {
			respond.Err(err)
			return
		}

		jwtPayload, err := a.Strategy.VerifyAccess(token)
		if err != nil {
			respond.Err(err)
			return
		}

		ctx := helper.UserSetCtx(r.Context(), jwtPayload)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

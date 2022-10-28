package middlewares

import (
	"net/http"

	"primedividend/api/internal/infrastructure/server/middlewares/helper"
	"primedividend/api/internal/infrastructure/server/response"
	"primedividend/api/internal/modules/auth/service/strategy"
)

type auth struct {
	strategy strategy.Strategy
}

func NewAuth(strategy strategy.Strategy) auth {
	return auth{strategy: strategy}
}

func (a auth) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		respond := response.NewRespondBuilder(w, r)

		token, err := helper.TokenFromRequest(r)
		if err != nil {
			respond.Err(err)
			return
		}

		jwtPayload, err := a.strategy.VerifyAccess(token)
		if err != nil {
			respond.Err(err)
			return
		}

		ctx := helper.UserSetCtx(r.Context(), jwtPayload)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

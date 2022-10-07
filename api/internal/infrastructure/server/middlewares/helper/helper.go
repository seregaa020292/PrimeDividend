package helper

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5/middleware"

	"primedivident/internal/config/consts"
	"primedivident/internal/modules/auth/entity"
	"primedivident/pkg/errs"
	"primedivident/pkg/errs/errmsg"
	"primedivident/pkg/logger"
)

func GetLogEntry(r *http.Request) logger.Logger {
	entry := middleware.GetLogEntry(r).(logger.Logger)
	return entry
}

func TokenFromRequest(r *http.Request) (string, error) {
	token := r.Header.Get("Authorization")
	if token == "" {
		token = r.FormValue("access_token")
	}

	if token = TokenPayload(token); token == "" {
		return "", errs.Unauthorized.New(errmsg.FailedGetData)
	}

	return token, nil
}

func TokenPayload(token string) string {
	if len(token) > 7 && strings.EqualFold(token[0:6], consts.TokenType) {
		return token[7:]
	}
	return ""
}

type userContextKey struct{}

func UserSetCtx(ctx context.Context, value entity.JwtPayload) context.Context {
	return context.WithValue(ctx, userContextKey{}, value)
}

func UserFromCtx(ctx context.Context) (entity.JwtPayload, error) {
	if jwtPayload, ok := ctx.Value(userContextKey{}).(entity.JwtPayload); ok {
		return jwtPayload, nil
	}

	return entity.JwtPayload{}, errs.Unauthorized.New(errmsg.AuthorizationRequired)
}

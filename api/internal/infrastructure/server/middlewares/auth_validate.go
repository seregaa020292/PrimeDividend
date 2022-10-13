package middlewares

import (
	"context"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers"

	"primedivident/internal/infrastructure/server/middlewares/helper"
	"primedivident/internal/infrastructure/server/response"
	"primedivident/internal/modules/auth/entity"
	"primedivident/internal/modules/auth/service/strategy"
	"primedivident/pkg/errs"
	"primedivident/pkg/errs/errmsg"
)

type authValidate struct {
	router   routers.Router
	strategy strategy.Strategy
}

func NewAuthValidate(router routers.Router, strategy strategy.Strategy) authValidate {
	return authValidate{
		router:   router,
		strategy: strategy,
	}
}

func (a authValidate) Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		respond := response.NewRespondBuilder(w, r)

		route, pathParams, err := a.router.FindRoute(r)
		if err != nil {
			respond.Err(err)
			return
		}

		var jwtPayload entity.JwtPayload

		requestValidationInput := &openapi3filter.RequestValidationInput{
			Request:    r,
			PathParams: pathParams,
			Route:      route,
			Options: &openapi3filter.Options{
				AuthenticationFunc: func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
					token, err := helper.TokenFromRequest(input.RequestValidationInput.Request)
					if err != nil {
						return err
					}

					jwtPayload, err = a.strategy.VerifyAccess(token)

					return err
				},
			},
		}

		err = openapi3filter.ValidateRequest(context.Background(), requestValidationInput)
		if err != nil {
			switch err.(type) {
			case *openapi3filter.RequestError:
				err = errs.BadRequest.Wrap(err, errmsg.AccessDenied)
			case *openapi3filter.SecurityRequirementsError:
				err = errs.Unauthorized.Wrap(err, errmsg.AuthorizationRequired)
			default:
				err = errs.ServerError.Wrap(err, errmsg.ServerError)
			}

			respond.Err(err)
			return
		}

		ctx := helper.UserSetCtx(r.Context(), jwtPayload)
		next(w, r.WithContext(ctx))
	}
}

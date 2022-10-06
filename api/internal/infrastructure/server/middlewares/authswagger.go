package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers"

	"primedivident/internal/infrastructure/server/response"
	"primedivident/pkg/errs"
	"primedivident/pkg/errs/errmsg"
)

func AuthSwagger(router routers.Router, verify func(token string) error) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			respond := response.NewRespondBuilder(w, r)

			route, pathParams, err := router.FindRoute(r)
			if err != nil {
				respond.Err(err)
				return
			}

			requestValidationInput := &openapi3filter.RequestValidationInput{
				Request:    r,
				PathParams: pathParams,
				Route:      route,
				Options: &openapi3filter.Options{
					AuthenticationFunc: func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
						bearerToken := input.RequestValidationInput.Request.Header.Get("Authorization")
						scheme := fmt.Sprintf("%s ", input.SecurityScheme.Scheme)
						accessToken := strings.Replace(bearerToken, scheme, "", 1)

						return verify(accessToken)
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

			next(w, r)
		}
	}
}

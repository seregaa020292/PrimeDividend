package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	middleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"

	"primedivident/internal/infrastructure/server/response"
	"primedivident/pkg/errs"
	"primedivident/pkg/errs/errmsg"
)

func AuthSwagger(swagger *openapi3.T, verify func(token string) error) func(next http.Handler) http.Handler {
	return middleware.OapiRequestValidatorWithOptions(swagger, &middleware.Options{
		Options: openapi3filter.Options{
			AuthenticationFunc: func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
				bearerToken := input.RequestValidationInput.Request.Header.Get("Authorization")
				scheme := fmt.Sprintf("%s ", input.SecurityScheme.Scheme)
				accessToken := strings.Replace(bearerToken, scheme, "", 1)

				return verify(accessToken)
			},
		},
		ErrorHandler: func(w http.ResponseWriter, message string, statusCode int) {
			respond := response.NewRespondBuilder(w, &http.Request{})
			err := errs.New(message)

			if statusCode == 400 {
				err = errs.NotFound.Wrap(err, errmsg.CouldNotBeFound)
			} else {
				err = errs.Forbidden.Wrap(err, errmsg.AccessDenied)
			}

			respond.Err(err)
		},
	})
}

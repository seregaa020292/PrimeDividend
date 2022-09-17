package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	middleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers"

	"primedivident/pkg/errorn"
	"primedivident/pkg/response"
)

func authValidator(swagger *openapi3.T) func(next http.Handler) http.Handler {
	return middleware.OapiRequestValidatorWithOptions(swagger, &middleware.Options{
		Options: openapi3filter.Options{
			AuthenticationFunc: func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
				return nil
			},
		},
		ErrorHandler: func(w http.ResponseWriter, message string, statusCode int) {
			respond := response.NewRespondBuilder(w, &http.Request{})

			var err error
			if statusCode == 400 {
				err = errorn.ErrorNotFoundElement
			} else {
				err = errorn.ErrorAccess.Wrap(fmt.Errorf("%s", message))
			}

			respond.Err(err)
		},
	})
}

func custom(router routers.Router) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			route, _, err := router.FindRoute(r)
			if err != nil {
				panic(err)
			}

			log.Print(route.Operation.OperationID)

			next(w, r)
		}
	}
}

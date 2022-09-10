package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	middleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"

	"primedivident/internal/infrastructure/http/openapi"
	"primedivident/pkg/response"
)

func newOpenapi() []middlewareFunc {
	swagger, err := openapi.GetSwagger()
	if err != nil {
		log.Fatalln(err)
	}

	swagger.Servers = nil

	return []middlewareFunc{
		authValidator(swagger),
	}
}

func authValidator(swagger *openapi3.T) middlewareFunc {
	return middleware.OapiRequestValidatorWithOptions(swagger, &middleware.Options{
		Options: openapi3filter.Options{
			AuthenticationFunc: func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
				return nil
			},
		},
		ErrorHandler: func(w http.ResponseWriter, message string, statusCode int) {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(statusCode)

			_ = json.NewEncoder(w).Encode(
				response.ErrRender(fmt.Errorf("%s", message)),
			)
		},
	})
}

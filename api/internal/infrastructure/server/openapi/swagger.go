package openapi

import (
	"log"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/routers"
	"github.com/getkin/kin-openapi/routers/gorillamux"
)

type Swagger struct {
	Swagger *openapi3.T
	Router  routers.Router
}

func NewSwagger() Swagger {
	swagger, err := GetSwagger()
	if err != nil {
		log.Fatalln(err)
	}

	swagger.Servers = nil

	router, err := gorillamux.NewRouter(swagger)
	if err != nil {
		log.Fatalln(err)
	}

	return Swagger{
		Swagger: swagger,
		Router:  router,
	}
}

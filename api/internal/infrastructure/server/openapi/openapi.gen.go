// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package openapi

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Создание пользователя
	// (POST /auth/email)
	JoinEmail(w http.ResponseWriter, r *http.Request)
	// Подтверждение почты
	// (POST /auth/email/confirm)
	ConfirmEmail(w http.ResponseWriter, r *http.Request)
	// Аутентификация по Email
	// (POST /auth/email/login)
	LoginEmail(w http.ResponseWriter, r *http.Request)
	// Выход из аккаунта
	// (POST /auth/logout)
	Logout(w http.ResponseWriter, r *http.Request)
	// Аутентификация в соц. сети
	// (GET /auth/{network})
	JoinNetwork(w http.ResponseWriter, r *http.Request, network Network)
	// Авторизация в соц. сети
	// (GET /auth/{network}/callback)
	ConfirmNetwork(w http.ResponseWriter, r *http.Request, network Network)

	// (GET /instrument)
	GetInstruments(w http.ResponseWriter, r *http.Request)
	// Создание портфеля
	// (POST /portfolio)
	CreatePortfolio(w http.ResponseWriter, r *http.Request)
	// Получение портфеля по ID
	// (GET /portfolio/{portfolioId})
	GetPortfolio(w http.ResponseWriter, r *http.Request, portfolioId PortfolioId)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// JoinEmail operation middleware
func (siw *ServerInterfaceWrapper) JoinEmail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.JoinEmail(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// ConfirmEmail operation middleware
func (siw *ServerInterfaceWrapper) ConfirmEmail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ConfirmEmail(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// LoginEmail operation middleware
func (siw *ServerInterfaceWrapper) LoginEmail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.LoginEmail(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// Logout operation middleware
func (siw *ServerInterfaceWrapper) Logout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.Logout(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// JoinNetwork operation middleware
func (siw *ServerInterfaceWrapper) JoinNetwork(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "network" -------------
	var network Network

	err = runtime.BindStyledParameter("simple", false, "network", chi.URLParam(r, "network"), &network)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "network", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.JoinNetwork(w, r, network)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// ConfirmNetwork operation middleware
func (siw *ServerInterfaceWrapper) ConfirmNetwork(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "network" -------------
	var network Network

	err = runtime.BindStyledParameter("simple", false, "network", chi.URLParam(r, "network"), &network)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "network", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ConfirmNetwork(w, r, network)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetInstruments operation middleware
func (siw *ServerInterfaceWrapper) GetInstruments(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetInstruments(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// CreatePortfolio operation middleware
func (siw *ServerInterfaceWrapper) CreatePortfolio(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreatePortfolio(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetPortfolio operation middleware
func (siw *ServerInterfaceWrapper) GetPortfolio(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "portfolioId" -------------
	var portfolioId PortfolioId

	err = runtime.BindStyledParameter("simple", false, "portfolioId", chi.URLParam(r, "portfolioId"), &portfolioId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "portfolioId", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetPortfolio(w, r, portfolioId)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/auth/email", wrapper.JoinEmail)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/auth/email/confirm", wrapper.ConfirmEmail)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/auth/email/login", wrapper.LoginEmail)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/auth/logout", wrapper.Logout)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/auth/{network}", wrapper.JoinNetwork)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/auth/{network}/callback", wrapper.ConfirmNetwork)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/instrument", wrapper.GetInstruments)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/portfolio", wrapper.CreatePortfolio)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/portfolio/{portfolioId}", wrapper.GetPortfolio)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"0FFnU5nQJEEQNlv/DgAA//+3iMKTGhwAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}

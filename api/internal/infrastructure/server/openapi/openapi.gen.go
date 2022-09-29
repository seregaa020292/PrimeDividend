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
	// Обновить токены
	// (POST /auth/refresh-token)
	RefreshToken(w http.ResponseWriter, r *http.Request)
	// Аутентификация через соц. сети
	// (GET /auth/{network})
	JoinNetwork(w http.ResponseWriter, r *http.Request, network Network)
	// Авторизация через соц. сети
	// (GET /auth/{network}/callback)
	ConfirmNetwork(w http.ResponseWriter, r *http.Request, network Network, params ConfirmNetworkParams)

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

// RefreshToken operation middleware
func (siw *ServerInterfaceWrapper) RefreshToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.RefreshToken(w, r)
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

	// Parameter object where we will unmarshal all parameters from the context
	var params ConfirmNetworkParams

	// ------------- Required query parameter "code" -------------
	if paramValue := r.URL.Query().Get("code"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "code"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "code", r.URL.Query(), &params.Code)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "code", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ConfirmNetwork(w, r, network, params)
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
		r.Post(options.BaseURL+"/auth/refresh-token", wrapper.RefreshToken)
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
	"5zKhUiIIG73/BgAA//+VClKuZx8AAA==",
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

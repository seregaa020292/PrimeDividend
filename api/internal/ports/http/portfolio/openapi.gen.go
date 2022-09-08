// Package portfolio provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package portfolio

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/go-chi/chi/v5"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Error defines model for error.
type Error struct {
	Data   *interface{} `json:"data"`
	Errors []struct {
		// Поле в котором произошла ошибка
		Field *string `json:"field,omitempty"`

		// Описание ошибки
		Message string `json:"message"`
	} `json:"errors"`
}

// Portfolio defines model for portfolio.
type Portfolio struct {
	CreatedAt time.Time          `json:"createdAt"`
	Id        openapi_types.UUID `json:"id"`
}

// PortfolioUpdate defines model for portfolioUpdate.
type PortfolioUpdate struct {
	CurrencyId openapi_types.UUID `json:"currencyId" validate:"required,uuid"`
	Title      string             `json:"title" validate:"required"`
	UserId     openapi_types.UUID `json:"userId" validate:"required,uuid"`
}

// PortfolioId defines model for portfolioId.
type PortfolioId = openapi_types.UUID

// N500 defines model for 500.
type N500 = Error

// CreatePortfolioJSONBody defines parameters for CreatePortfolio.
type CreatePortfolioJSONBody = PortfolioUpdate

// CreatePortfolioJSONRequestBody defines body for CreatePortfolio for application/json ContentType.
type CreatePortfolioJSONRequestBody = CreatePortfolioJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Создание портфеля
	// (POST /portfolio)
	CreatePortfolio(w http.ResponseWriter, r *http.Request)
	// Получение портфеля по ID
	// (GET /portfolio/{portfolioId})
	GetPortfolioById(w http.ResponseWriter, r *http.Request, portfolioId PortfolioId)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

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

// GetPortfolioById operation middleware
func (siw *ServerInterfaceWrapper) GetPortfolioById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "portfolioId" -------------
	var portfolioId PortfolioId

	err = runtime.BindStyledParameter("simple", false, "portfolioId", chi.URLParam(r, "portfolioId"), &portfolioId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "portfolioId", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetPortfolioById(w, r, portfolioId)
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
		r.Post(options.BaseURL+"/portfolio", wrapper.CreatePortfolio)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/portfolio/{portfolioId}", wrapper.GetPortfolioById)
	})

	return r
}

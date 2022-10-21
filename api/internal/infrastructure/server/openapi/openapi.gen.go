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
	// Получение всех валют
	// (GET /currency)
	GetCurrencies(w http.ResponseWriter, r *http.Request)
	// Получение валюты по ID
	// (GET /currency/{currencyId})
	GetCurrency(w http.ResponseWriter, r *http.Request, currencyId CurrencyId)
	// Получение инструментов
	// (GET /instrument)
	GetInstruments(w http.ResponseWriter, r *http.Request)
	// Получение инструмента по ID
	// (GET /instrument/{instrumentId})
	GetInstrument(w http.ResponseWriter, r *http.Request, instrumentId InstrumentId)
	// Получение списка ценных бумаг
	// (GET /market)
	GetMarkets(w http.ResponseWriter, r *http.Request, params GetMarketsParams)
	// Получение ценной бумаги по ID
	// (GET /market/{marketId})
	GetMarket(w http.ResponseWriter, r *http.Request, marketId MarketId)
	// Получение ценной бумаги по тикеру
	// (GET /market/{ticker})
	GetMarketByTicker(w http.ResponseWriter, r *http.Request, ticker Ticker)
	// Получение списка портфелей
	// (GET /portfolio)
	GetPortfolios(w http.ResponseWriter, r *http.Request, params GetPortfoliosParams)
	// Создание портфеля
	// (POST /portfolio)
	CreatePortfolio(w http.ResponseWriter, r *http.Request)
	// Удаление портфеля
	// (DELETE /portfolio/{portfolioId})
	RemovePortfolio(w http.ResponseWriter, r *http.Request, portfolioId PortfolioId)
	// Получение портфеля по ID
	// (GET /portfolio/{portfolioId})
	GetPortfolio(w http.ResponseWriter, r *http.Request, portfolioId PortfolioId)
	// Редактирование портфеля
	// (PATCH /portfolio/{portfolioId})
	UpdatePortfolio(w http.ResponseWriter, r *http.Request, portfolioId PortfolioId)
	// Получение поставщиков котировок
	// (GET /provider)
	GetProviders(w http.ResponseWriter, r *http.Request)
	// Получение поставщика котировок по ID
	// (GET /provider/{providerId})
	GetProvider(w http.ResponseWriter, r *http.Request, providerId ProviderId)
	// Получение данных пользователя
	// (GET /user)
	GetUser(w http.ResponseWriter, r *http.Request)
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

	// ------------- Required query parameter "state" -------------
	if paramValue := r.URL.Query().Get("state"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "state"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "state", r.URL.Query(), &params.State)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "state", Err: err})
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

// GetCurrencies operation middleware
func (siw *ServerInterfaceWrapper) GetCurrencies(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetCurrencies(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetCurrency operation middleware
func (siw *ServerInterfaceWrapper) GetCurrency(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "currencyId" -------------
	var currencyId CurrencyId

	err = runtime.BindStyledParameter("simple", false, "currencyId", chi.URLParam(r, "currencyId"), &currencyId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "currencyId", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetCurrency(w, r, currencyId)
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

// GetInstrument operation middleware
func (siw *ServerInterfaceWrapper) GetInstrument(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "instrumentId" -------------
	var instrumentId InstrumentId

	err = runtime.BindStyledParameter("simple", false, "instrumentId", chi.URLParam(r, "instrumentId"), &instrumentId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "instrumentId", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetInstrument(w, r, instrumentId)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetMarkets operation middleware
func (siw *ServerInterfaceWrapper) GetMarkets(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetMarketsParams

	// ------------- Optional query parameter "limit" -------------
	if paramValue := r.URL.Query().Get("limit"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "limit", Err: err})
		return
	}

	// ------------- Optional query parameter "cursor" -------------
	if paramValue := r.URL.Query().Get("cursor"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "cursor", r.URL.Query(), &params.Cursor)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "cursor", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetMarkets(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetMarket operation middleware
func (siw *ServerInterfaceWrapper) GetMarket(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "marketId" -------------
	var marketId MarketId

	err = runtime.BindStyledParameter("simple", false, "marketId", chi.URLParam(r, "marketId"), &marketId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "marketId", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetMarket(w, r, marketId)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetMarketByTicker operation middleware
func (siw *ServerInterfaceWrapper) GetMarketByTicker(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "ticker" -------------
	var ticker Ticker

	err = runtime.BindStyledParameter("simple", false, "ticker", chi.URLParam(r, "ticker"), &ticker)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "ticker", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetMarketByTicker(w, r, ticker)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetPortfolios operation middleware
func (siw *ServerInterfaceWrapper) GetPortfolios(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPortfoliosParams

	// ------------- Optional query parameter "limit" -------------
	if paramValue := r.URL.Query().Get("limit"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "limit", Err: err})
		return
	}

	// ------------- Optional query parameter "cursor" -------------
	if paramValue := r.URL.Query().Get("cursor"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "cursor", r.URL.Query(), &params.Cursor)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "cursor", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetPortfolios(w, r, params)
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

// RemovePortfolio operation middleware
func (siw *ServerInterfaceWrapper) RemovePortfolio(w http.ResponseWriter, r *http.Request) {
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
		siw.Handler.RemovePortfolio(w, r, portfolioId)
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

// UpdatePortfolio operation middleware
func (siw *ServerInterfaceWrapper) UpdatePortfolio(w http.ResponseWriter, r *http.Request) {
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
		siw.Handler.UpdatePortfolio(w, r, portfolioId)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetProviders operation middleware
func (siw *ServerInterfaceWrapper) GetProviders(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetProviders(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetProvider operation middleware
func (siw *ServerInterfaceWrapper) GetProvider(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "providerId" -------------
	var providerId ProviderId

	err = runtime.BindStyledParameter("simple", false, "providerId", chi.URLParam(r, "providerId"), &providerId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "providerId", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetProvider(w, r, providerId)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetUser operation middleware
func (siw *ServerInterfaceWrapper) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUser(w, r)
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
		r.Get(options.BaseURL+"/currency", wrapper.GetCurrencies)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/currency/{currencyId}", wrapper.GetCurrency)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/instrument", wrapper.GetInstruments)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/instrument/{instrumentId}", wrapper.GetInstrument)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/market", wrapper.GetMarkets)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/market/{marketId}", wrapper.GetMarket)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/market/{ticker}", wrapper.GetMarketByTicker)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/portfolio", wrapper.GetPortfolios)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/portfolio", wrapper.CreatePortfolio)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/portfolio/{portfolioId}", wrapper.RemovePortfolio)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/portfolio/{portfolioId}", wrapper.GetPortfolio)
	})
	r.Group(func(r chi.Router) {
		r.Patch(options.BaseURL+"/portfolio/{portfolioId}", wrapper.UpdatePortfolio)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/provider", wrapper.GetProviders)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/provider/{providerId}", wrapper.GetProvider)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/user", wrapper.GetUser)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
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

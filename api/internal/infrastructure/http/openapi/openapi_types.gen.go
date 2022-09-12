// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package openapi

import (
	"time"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

const (
	BearerAuthScopes = "BearerAuth.Scopes"
)

// AuthConfirm defines model for authConfirm.
type AuthConfirm struct {
	Token string `json:"token"`
}

// AuthUser defines model for authUser.
type AuthUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

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

// Instrument defines model for instrument.
type Instrument struct {
	CreatedAt   time.Time          `json:"createdAt"`
	Description string             `json:"description"`
	Id          openapi_types.UUID `json:"id"`
	Title       string             `json:"title"`
	UpdatedAt   time.Time          `json:"updatedAt"`
}

// Instruments defines model for instruments.
type Instruments = []Instrument

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

// Network defines model for network.
type Network = string

// PortfolioId defines model for portfolioId.
type PortfolioId = openapi_types.UUID

// N400 defines model for 400.
type N400 = Error

// N500 defines model for 500.
type N500 = Error

// AuthEmailJSONBody defines parameters for AuthEmail.
type AuthEmailJSONBody = AuthUser

// AuthEmailConfirmJSONBody defines parameters for AuthEmailConfirm.
type AuthEmailConfirmJSONBody = AuthConfirm

// CreatePortfolioJSONBody defines parameters for CreatePortfolio.
type CreatePortfolioJSONBody = PortfolioUpdate

// AuthEmailJSONRequestBody defines body for AuthEmail for application/json ContentType.
type AuthEmailJSONRequestBody = AuthEmailJSONBody

// AuthEmailConfirmJSONRequestBody defines body for AuthEmailConfirm for application/json ContentType.
type AuthEmailConfirmJSONRequestBody = AuthEmailConfirmJSONBody

// CreatePortfolioJSONRequestBody defines body for CreatePortfolio for application/json ContentType.
type CreatePortfolioJSONRequestBody = CreatePortfolioJSONBody

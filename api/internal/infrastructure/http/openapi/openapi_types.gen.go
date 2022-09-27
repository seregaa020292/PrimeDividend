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
	Token openapi_types.UUID `json:"token" validate:"required"`
}

// AuthUser defines model for authUser.
type AuthUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// Error defines model for error.
type Error struct {
	Data *interface{} `json:"data"`

	// Объект ошибки
	Error *struct {
		// В каком поле произошла ошибка
		Details *[]struct {
			// Описание ошибки
			Message string `json:"message"`

			// Поле в котором произошла ошибка
			Target *string `json:"target,omitempty"`
		} `json:"details,omitempty"`

		// Описание ошибки
		Message string `json:"message"`

		// HTTP код ошибки
		Status float32 `json:"status"`

		// Место где произошла ошибка
		Target float32 `json:"target"`
	} `json:"error,omitempty"`
}

// Instrument defines model for instrument.
type Instrument struct {
	Description string             `json:"description"`
	Id          openapi_types.UUID `json:"id"`
	Title       string             `json:"title"`
}

// Instruments defines model for instruments.
type Instruments = []Instrument

// LoginUser defines model for loginUser.
type LoginUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
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

// Network defines model for network.
type Network = string

// PortfolioId defines model for portfolioId.
type PortfolioId = openapi_types.UUID

// N400 defines model for 400.
type N400 = Error

// N403 defines model for 403.
type N403 = Error

// N500 defines model for 500.
type N500 = Error

// JoinEmailJSONBody defines parameters for JoinEmail.
type JoinEmailJSONBody = AuthUser

// ConfirmEmailJSONBody defines parameters for ConfirmEmail.
type ConfirmEmailJSONBody = AuthConfirm

// LoginEmailJSONBody defines parameters for LoginEmail.
type LoginEmailJSONBody = LoginUser

// CreatePortfolioJSONBody defines parameters for CreatePortfolio.
type CreatePortfolioJSONBody = PortfolioUpdate

// JoinEmailJSONRequestBody defines body for JoinEmail for application/json ContentType.
type JoinEmailJSONRequestBody = JoinEmailJSONBody

// ConfirmEmailJSONRequestBody defines body for ConfirmEmail for application/json ContentType.
type ConfirmEmailJSONRequestBody = ConfirmEmailJSONBody

// LoginEmailJSONRequestBody defines body for LoginEmail for application/json ContentType.
type LoginEmailJSONRequestBody = LoginEmailJSONBody

// CreatePortfolioJSONRequestBody defines body for CreatePortfolio for application/json ContentType.
type CreatePortfolioJSONRequestBody = CreatePortfolioJSONBody

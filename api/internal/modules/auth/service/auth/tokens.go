package auth

import (
	"primedivident/internal/config"
	"primedivident/pkg/token"
)

type JwtTokens interface {
	GenTokens(data any) (Tokens, error)
	GenAccessToken(data any) (string, error)
	GenRefreshToken() (string, error)
}

type jwtTokens struct {
	accessTokenService  token.JwtService
	refreshTokenService token.JwtService
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewJwtTokens(issuer string, jwt config.Jwt) JwtTokens {
	return jwtTokens{
		accessTokenService:  token.NewJwtService(issuer, jwt.AccessSecretKey, jwt.AccessExpiresIn),
		refreshTokenService: token.NewJwtService(issuer, jwt.RefreshSecretKey, jwt.RefreshExpiresIn),
	}
}

func (t jwtTokens) GenAccessToken(data any) (string, error) {
	return t.accessTokenService.GenerateToken(data)
}

func (t jwtTokens) GenRefreshToken() (string, error) {
	return t.refreshTokenService.GenerateToken(struct{}{})
}

func (t jwtTokens) GenTokens(data any) (Tokens, error) {
	var (
		tokens Tokens
		err    error
	)

	if tokens.AccessToken, err = t.GenAccessToken(data); err != nil {
		return Tokens{}, err
	}

	if tokens.RefreshToken, err = t.GenRefreshToken(); err != nil {
		return Tokens{}, err
	}

	return tokens, nil
}

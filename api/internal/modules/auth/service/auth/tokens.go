package auth

import (
	"primedivident/internal/config"
	"primedivident/internal/modules/auth/entity"
	"primedivident/pkg/token"
)

type JwtTokens interface {
	GenTokens(entity.JwtUser) (Tokens, error)
	GenAccessToken(entity.JwtUser) (string, error)
	ValidateAccessToken(token string) (entity.JwtUser, error)
	GenRefreshToken() (string, error)
}

type jwtTokens struct {
	accessTokenService  token.JwtService[entity.JwtUser]
	refreshTokenService token.JwtService[any]
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewJwtTokens(issuer string, jwt config.Jwt) JwtTokens {
	return jwtTokens{
		accessTokenService:  token.NewJwtService[entity.JwtUser](issuer, jwt.AccessSecretKey, jwt.AccessExpiresIn),
		refreshTokenService: token.NewJwtService[any](issuer, jwt.RefreshSecretKey, jwt.RefreshExpiresIn),
	}
}

func (t jwtTokens) GenAccessToken(data entity.JwtUser) (string, error) {
	return t.accessTokenService.GenerateToken(&data)
}

func (t jwtTokens) ValidateAccessToken(token string) (entity.JwtUser, error) {
	data, err := t.accessTokenService.ValidateToken(token)
	if err != nil {
		return entity.JwtUser{}, err
	}

	return *data, nil
}

func (t jwtTokens) GenRefreshToken() (string, error) {
	return t.refreshTokenService.GenerateToken(nil)
}

func (t jwtTokens) GenTokens(data entity.JwtUser) (Tokens, error) {
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

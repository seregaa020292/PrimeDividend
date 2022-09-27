package entity

import (
	"github.com/google/uuid"

	"primedivident/internal/config"
	"primedivident/pkg/token"
)

type JwtTokens interface {
	GenTokens(JwtUser) (Tokens, error)
	GenAccessToken(JwtUser) (token.Token, error)
	ValidateAccessToken(token string) (JwtUser, error)
	GenRefreshToken() (token.Token, error)
}

type Tokens struct {
	AccessToken  token.Token
	RefreshToken token.Token
}

type jwtTokens struct {
	accessTokenService  token.JwtService[JwtUser]
	refreshTokenService token.JwtService[uuid.UUID]
}

func NewJwtTokens(issuer string, jwt config.Jwt) JwtTokens {
	return jwtTokens{
		accessTokenService:  token.NewJwtService[JwtUser](issuer, jwt.AccessSecretKey, jwt.AccessExpiresIn),
		refreshTokenService: token.NewJwtService[uuid.UUID](issuer, jwt.RefreshSecretKey, jwt.RefreshExpiresIn),
	}
}

func (t jwtTokens) GenAccessToken(data JwtUser) (token.Token, error) {
	return t.accessTokenService.GenerateToken(data)
}

func (t jwtTokens) ValidateAccessToken(token string) (JwtUser, error) {
	data, err := t.accessTokenService.ValidateToken(token)
	if err != nil {
		return JwtUser{}, err
	}

	return data, nil
}

func (t jwtTokens) GenRefreshToken() (token.Token, error) {
	return t.refreshTokenService.GenerateToken(uuid.New())
}

func (t jwtTokens) GenTokens(data JwtUser) (Tokens, error) {
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

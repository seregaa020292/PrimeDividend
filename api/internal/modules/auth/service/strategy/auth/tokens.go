package auth

import (
	"github.com/google/uuid"

	"primedivident/internal/config"
	"primedivident/internal/modules/auth/entity"
	"primedivident/pkg/token"
)

type JwtTokens interface {
	GenTokens(entity.JwtUser) (Tokens, error)
	GenAccessToken(entity.JwtUser) (token.Token, error)
	ValidateAccessToken(token string) (entity.JwtUser, error)
	GenRefreshToken() (token.Token, error)
}

type Tokens struct {
	AccessToken  token.Token
	RefreshToken token.Token
}

type jwtTokens struct {
	accessTokenService  token.JwtService[entity.JwtUser]
	refreshTokenService token.JwtService[uuid.UUID]
}

func NewJwtTokens(issuer string, jwt config.Jwt) JwtTokens {
	return jwtTokens{
		accessTokenService:  token.NewJwtService[entity.JwtUser](issuer, jwt.AccessSecretKey, jwt.AccessExpiresIn),
		refreshTokenService: token.NewJwtService[uuid.UUID](issuer, jwt.RefreshSecretKey, jwt.RefreshExpiresIn),
	}
}

func (t jwtTokens) GenAccessToken(data entity.JwtUser) (token.Token, error) {
	return t.accessTokenService.GenerateToken(data)
}

func (t jwtTokens) ValidateAccessToken(token string) (entity.JwtUser, error) {
	data, err := t.accessTokenService.ValidateToken(token)
	if err != nil {
		return entity.JwtUser{}, err
	}

	return data, nil
}

func (t jwtTokens) GenRefreshToken() (token.Token, error) {
	return t.refreshTokenService.GenerateToken(uuid.New())
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

package auth

import (
	"github.com/google/uuid"

	"primedivident/internal/config"
	"primedivident/internal/modules/auth/entity"
	"primedivident/pkg/token"
	"primedivident/pkg/utils/gog"
)

type JwtTokens interface {
	GenTokens(entity.JwtUser) (Tokens, error)
	GenAccessToken(entity.JwtUser) (string, error)
	ValidateAccessToken(token string) (entity.JwtUser, error)
	GenRefreshToken() (string, error)
}

type jwtTokens struct {
	accessTokenService  token.JwtService[entity.JwtUser]
	refreshTokenService token.JwtService[uuid.UUID]
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewJwtTokens(issuer string, jwt config.Jwt) JwtTokens {
	return jwtTokens{
		accessTokenService:  token.NewJwtService[entity.JwtUser](issuer, jwt.AccessSecretKey, jwt.AccessExpiresIn),
		refreshTokenService: token.NewJwtService[uuid.UUID](issuer, jwt.RefreshSecretKey, jwt.RefreshExpiresIn),
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
	return t.refreshTokenService.GenerateToken(gog.Ptr(uuid.New()))
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

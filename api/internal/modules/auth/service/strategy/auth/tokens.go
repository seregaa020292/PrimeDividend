package auth

import (
	"github.com/google/uuid"

	"primedivident/internal/config"
	"primedivident/internal/modules/auth/entity"
	"primedivident/pkg/token"
)

type JwtTokens interface {
	GenTokens(entity.JwtPayload) (Tokens, error)
	GenAccessToken(entity.JwtPayload) (token.Token, error)
	ValidateAccessToken(token string) (entity.JwtPayload, error)
	GenRefreshToken() (token.Token, error)
}

type Tokens struct {
	AccessToken  token.Token
	RefreshToken token.Token
}

type jwtTokens struct {
	accessTokenService  token.JwtService[entity.JwtPayload]
	refreshTokenService token.JwtService[uuid.UUID]
}

func NewJwtTokens(issuer string, cfg config.Jwt) JwtTokens {
	return jwtTokens{
		accessTokenService:  token.NewJwtService[entity.JwtPayload](issuer, cfg.AccessSecretKey, cfg.AccessExpiresIn),
		refreshTokenService: token.NewJwtService[uuid.UUID](issuer, cfg.RefreshSecretKey, cfg.RefreshExpiresIn),
	}
}

func (t jwtTokens) GenAccessToken(data entity.JwtPayload) (token.Token, error) {
	return t.accessTokenService.GenerateToken(data)
}

func (t jwtTokens) ValidateAccessToken(token string) (entity.JwtPayload, error) {
	data, err := t.accessTokenService.ValidateToken(token)
	if err != nil {
		return entity.JwtPayload{}, err
	}

	return data, nil
}

func (t jwtTokens) GenRefreshToken() (token.Token, error) {
	return t.refreshTokenService.GenerateToken(uuid.New())
}

func (t jwtTokens) GenTokens(data entity.JwtPayload) (Tokens, error) {
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

package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"primedividend/api/internal/config"
	"primedividend/api/internal/modules/auth/entity"
	"primedividend/api/pkg/token"
)

type JwtTokens interface {
	GenTokens(entity.JwtPayload) (Tokens, error)
	GenAccessToken(entity.JwtPayload) (token.Token, error)
	GenRefreshToken(id uuid.UUID) (token.Token, error)
	ValidateAccessToken(token string) (entity.JwtPayload, error)
	ValidateRefreshToken(token string) (uuid.UUID, error)
	CorrectRefreshToken(token string) error
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

func (t jwtTokens) GenRefreshToken(id uuid.UUID) (token.Token, error) {
	return t.refreshTokenService.GenerateToken(id)
}

func (t jwtTokens) ValidateAccessToken(token string) (entity.JwtPayload, error) {
	return t.accessTokenService.ValidateToken(token)
}

func (t jwtTokens) ValidateRefreshToken(token string) (uuid.UUID, error) {
	return t.refreshTokenService.ValidateToken(token)
}

func (t jwtTokens) CorrectRefreshToken(token string) error {
	if _, err := t.ValidateRefreshToken(token); err != nil {
		if !errors.Is(err, jwt.ErrTokenExpired) {
			return err
		}
	}

	return nil
}

func (t jwtTokens) GenTokens(data entity.JwtPayload) (Tokens, error) {
	var (
		tokens Tokens
		err    error
	)

	if tokens.AccessToken, err = t.GenAccessToken(data); err != nil {
		return Tokens{}, err
	}

	if tokens.RefreshToken, err = t.GenRefreshToken(data.ID); err != nil {
		return Tokens{}, err
	}

	return tokens, nil
}

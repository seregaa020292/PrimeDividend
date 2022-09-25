package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtService interface {
	GenerateToken(payload any) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
	issuer    string
	secretKey []byte
	expiresIn time.Duration
}

type jwtCustomClaims struct {
	Data any `json:"data"`
	jwt.RegisteredClaims
}

func NewJwtService(issuer string, secretKey string, expiresIn time.Duration) JwtService {
	return &jwtService{
		issuer:    issuer,
		secretKey: []byte(secretKey),
		expiresIn: expiresIn,
	}
}

func (j *jwtService) GenerateToken(data any) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwtCustomClaims{
		data,
		jwt.RegisteredClaims{
			Issuer:    j.issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.expiresIn)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})

	return token.SignedString(j.secretKey)
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return j.secretKey, nil
	})
}

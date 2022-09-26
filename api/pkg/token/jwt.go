package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type (
	Token struct {
		Value     string
		ExpiresAt time.Time
	}
	JwtService[Data any] interface {
		GenerateToken(data Data) (Token, error)
		ValidateToken(token string) (Data, error)
	}
)

type (
	jwtService[Data any] struct {
		issuer    string
		secretKey []byte
		expiresIn time.Duration
	}
	JwtCustomClaims[Data any] struct {
		Data Data `json:"data,omitempty"`
		jwt.RegisteredClaims
	}
)

func NewJwtService[Data any](issuer string, secretKey string, expiresIn time.Duration) JwtService[Data] {
	return &jwtService[Data]{
		issuer:    issuer,
		secretKey: []byte(secretKey),
		expiresIn: expiresIn,
	}
}

func (j *jwtService[Data]) GenerateToken(data Data) (Token, error) {
	nowAt := time.Now()
	expiresAt := nowAt.Add(j.expiresIn)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &JwtCustomClaims[Data]{
		Data: data,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.issuer,
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(nowAt),
		},
	})

	value, err := token.SignedString(j.secretKey)
	if err != nil {
		return Token{}, err
	}

	return Token{
		Value:     value,
		ExpiresAt: expiresAt,
	}, nil
}

func (j *jwtService[Data]) ValidateToken(token string) (Data, error) {
	var data Data

	verifySign := func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return data, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return j.secretKey, nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &JwtCustomClaims[Data]{}, verifySign)
	if err != nil {
		return data, err
	}

	claims, ok := jwtToken.Claims.(*JwtCustomClaims[Data])

	if ok && jwtToken.Valid {
		return claims.Data, nil
	}

	return data, errors.New("error validate token")
}

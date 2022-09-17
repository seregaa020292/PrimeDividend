package entity

import (
	"primedivident/internal/config/consts"
	"primedivident/pkg/utils/hash"
)

type User struct {
	Email    string
	PassHash string
	Token    Token
}

func NewUser(email, password string) (User, error) {
	pass, err := hash.Password(password).Hash()
	if err != nil {
		return User{}, err
	}

	return User{
		Email:    email,
		PassHash: pass,
		Token:    NewToken(consts.TokenJoinTTL),
	}, nil
}

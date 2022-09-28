package dto

import (
	"primedivident/internal/models/app/public/model"
	"primedivident/internal/modules/auth/entity"
)

func EntityUserByModel(user model.Users) entity.User {
	var token entity.Token
	if user.TokenJoinValue != nil {
		token.Value = *user.TokenJoinValue
	}
	if user.TokenJoinExpires != nil {
		token.Expires = *user.TokenJoinExpires
	}

	return entity.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		PassHash: user.Password,
		Role:     entity.Role(user.Role),
		Status:   entity.Status(user.Status),
		Token:    token,
	}
}

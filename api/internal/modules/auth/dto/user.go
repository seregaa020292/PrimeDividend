package dto

import (
	"primedivident/internal/models/app/public/model"
	"primedivident/internal/modules/auth/entity"
	"primedivident/pkg/utils/gog"
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

func ModelUserByEntity(user entity.User) model.Users {
	return model.Users{
		Email:            user.Email,
		Name:             user.Name,
		Password:         user.PassHash,
		Role:             user.Role.String(),
		Status:           user.Status.String(),
		TokenJoinValue:   gog.Ptr(user.Token.Value),
		TokenJoinExpires: gog.Ptr(user.Token.Expires),
	}
}

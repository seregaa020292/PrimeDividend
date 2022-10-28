package dto

import (
	"primedividend/api/internal/models"
	"primedividend/api/internal/models/app/public/model"
	"primedividend/api/internal/modules/auth/entity"
)

func EntityUserByModel(user model.Users) entity.User {
	return entity.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		PassHash: user.Password,
		Role:     entity.Role(user.Role),
		Status:   models.Status(user.Status),
		Token: entity.Token{
			Value:   user.TokenJoinValue,
			Expires: user.TokenJoinExpires,
		},
	}
}

func ModelUserByEntity(user entity.User) model.Users {
	return model.Users{
		ID:               user.ID,
		Email:            user.Email,
		Name:             user.Name,
		Password:         user.PassHash,
		Role:             user.Role.String(),
		Status:           user.Status.String(),
		TokenJoinValue:   user.Token.Value,
		TokenJoinExpires: user.Token.Expires,
	}
}

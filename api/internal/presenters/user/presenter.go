package user

import (
	"primedivident/internal/infrastructure/server/openapi"
	"primedivident/internal/models/app/public/model"
)

type Presenter interface {
	GetOne(portfolio model.Users) openapi.User
}

type present struct{}

func NewPresenter() Presenter {
	return present{}
}

func (p present) GetOne(item model.Users) openapi.User {
	return openapi.User{
		Id:        item.ID,
		Name:      item.Name,
		Email:     item.Email,
		Role:      item.Role,
		Avatar:    item.Avatar,
		Status:    item.Status,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}
}

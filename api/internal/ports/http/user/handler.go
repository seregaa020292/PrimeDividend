package user

import (
	"primedivident/internal/infrastructure/server/response"
	"primedivident/internal/modules/user/query"
	"primedivident/internal/presenters/user"
)

type HandlerUser struct {
	responder    response.Responder
	presenter    user.Presenter
	queryGetById query.GetById
}

func NewHandler(
	responder response.Responder,
	presenter user.Presenter,
	queryGetById query.GetById,
) HandlerUser {
	return HandlerUser{
		responder:    responder,
		presenter:    presenter,
		queryGetById: queryGetById,
	}
}

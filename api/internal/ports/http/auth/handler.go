package auth

import (
	"primedivident/internal/modules/auth/interactor/command"
	"primedivident/pkg/response"
)

type HandlerAuth struct {
	responder      response.Responder
	cmdJoinByEmail command.JoinByEmail
}

func NewHandler(
	responder response.Responder,
	cmdJoinByEmail command.JoinByEmail,
) HandlerAuth {
	return HandlerAuth{
		responder:      responder,
		cmdJoinByEmail: cmdJoinByEmail,
	}
}

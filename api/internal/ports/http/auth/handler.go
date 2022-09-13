package auth

import (
	"primedivident/internal/modules/auth/command"
	"primedivident/pkg/response"
)

type HandlerAuth struct {
	responder         response.Responder
	cmdJoinByEmail    command.JoinByEmail
	cmdConfirmByToken command.ConfirmByToken
}

func NewHandler(
	responder response.Responder,
	cmdJoinByEmail command.JoinByEmail,
	cmdConfirmByToken command.ConfirmByToken,
) HandlerAuth {
	return HandlerAuth{
		responder:         responder,
		cmdJoinByEmail:    cmdJoinByEmail,
		cmdConfirmByToken: cmdConfirmByToken,
	}
}

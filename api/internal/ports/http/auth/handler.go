package auth

import (
	"primedivident/internal/modules/auth/command"
	"primedivident/internal/modules/auth/service/auth"
	"primedivident/pkg/response"
)

type HandlerAuth struct {
	responder         response.Responder
	cmdJoinByEmail    command.JoinByEmail
	cmdConfirmByToken command.ConfirmByToken
	authService       auth.Auth
}

func NewHandler(
	responder response.Responder,
	cmdJoinByEmail command.JoinByEmail,
	cmdConfirmByToken command.ConfirmByToken,
	authService auth.Auth,
) HandlerAuth {
	return HandlerAuth{
		responder:         responder,
		cmdJoinByEmail:    cmdJoinByEmail,
		cmdConfirmByToken: cmdConfirmByToken,
		authService:       authService,
	}
}

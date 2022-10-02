package auth

import (
	"primedivident/internal/infrastructure/response"
	"primedivident/internal/modules/auth/command"
	"primedivident/internal/modules/auth/service/strategy"
)

type HandlerAuth struct {
	responder         response.Responder
	strategy          strategy.Strategy
	cmdJoinByEmail    command.JoinByEmail
	cmdConfirmByToken command.ConfirmByToken
}

func NewHandler(
	responder response.Responder,
	strategy strategy.Strategy,
	cmdJoinByEmail command.JoinByEmail,
	cmdConfirmByToken command.ConfirmByToken,
) HandlerAuth {
	return HandlerAuth{
		responder:         responder,
		strategy:          strategy,
		cmdJoinByEmail:    cmdJoinByEmail,
		cmdConfirmByToken: cmdConfirmByToken,
	}
}

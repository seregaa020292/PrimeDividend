package auth

import (
	"primedividend/api/internal/infrastructure/server/response"
	"primedividend/api/internal/modules/auth/command"
	"primedividend/api/internal/modules/auth/service/strategy"
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

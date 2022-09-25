package auth

import (
	"primedivident/internal/modules/auth/command"
	"primedivident/internal/modules/auth/service/auth/strategies"
	"primedivident/pkg/response"
)

type HandlerAuth struct {
	responder         response.Responder
	cmdJoinByEmail    command.JoinByEmail
	cmdConfirmByToken command.ConfirmByToken
	strategies        strategies.Strategies
}

func NewHandler(
	responder response.Responder,
	cmdJoinByEmail command.JoinByEmail,
	cmdConfirmByToken command.ConfirmByToken,
	strategies strategies.Strategies,
) HandlerAuth {
	return HandlerAuth{
		responder:         responder,
		cmdJoinByEmail:    cmdJoinByEmail,
		cmdConfirmByToken: cmdConfirmByToken,
		strategies:        strategies,
	}
}

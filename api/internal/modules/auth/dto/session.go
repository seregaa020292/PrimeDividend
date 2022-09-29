package dto

import (
	"github.com/google/uuid"

	"primedivident/internal/models/app/public/model"
	"primedivident/internal/modules/auth/entity"
	"primedivident/internal/modules/auth/service/strategy/auth"
	"primedivident/pkg/token"
)

func ModelSessionCreating(
	userID uuid.UUID,
	strategy auth.Name,
	refreshToken token.Token,
	accountability entity.Accountability,
) model.Sessions {
	return model.Sessions{
		Token:     refreshToken.Value,
		ExpiresAt: refreshToken.ExpiresAt,
		UserID:    userID,
		Strategy:  strategy.String(),
		IP:        accountability.IP,
		UserAgent: accountability.UserAgent,
		Origin:    accountability.Origin,
	}
}

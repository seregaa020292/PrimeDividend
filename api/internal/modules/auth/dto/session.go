package dto

import (
	"github.com/google/uuid"

	"primedividend/api/internal/models/app/public/model"
	"primedividend/api/internal/modules/auth/entity"
	"primedividend/api/internal/modules/auth/service/strategy/auth"
	"primedividend/api/pkg/token"
)

func ModelSessionCreating(
	userID uuid.UUID,
	refreshToken token.Token,
	accountability auth.Accountability,
) model.Sessions {
	return model.Sessions{
		Token:     refreshToken.Value,
		ExpiresAt: refreshToken.ExpiresAt,
		UserID:    userID,
		IP:        accountability.IP,
		UserAgent: accountability.UserAgent,
		Origin:    accountability.Origin,
	}
}

func EntitySessionByModel(session model.Sessions) entity.Session {
	return entity.Session{
		ID:      session.ID,
		Token:   session.Token,
		Expires: session.ExpiresAt,
	}
}

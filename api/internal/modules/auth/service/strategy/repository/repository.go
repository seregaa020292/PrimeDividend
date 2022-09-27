package repository

import (
	jet "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"primedivident/internal/models/app/public/model"
	"primedivident/internal/models/app/public/table"
	"primedivident/internal/modules/auth/dto"
	"primedivident/internal/modules/auth/entity"
	"primedivident/internal/modules/auth/service/strategy/auth"
	"primedivident/pkg/db/postgres"
	"primedivident/pkg/token"
)

type Repository interface {
	FindUserByEmail(email string) (entity.User, error)
	FindUserByNetworkId(id string) (entity.User, error)
	AttachNetwork(user entity.JwtUser, network auth.Name)
	SaveRefreshToken(id uuid.UUID, refreshToken token.Token)
}

type repository struct {
	db *postgres.Postgres
}

func NewRepository(db *postgres.Postgres) Repository {
	return repository{db: db}
}

func (r repository) FindUserByEmail(email string) (entity.User, error) {
	var user model.Users

	stmt := table.Users.
		SELECT(table.Users.AllColumns).
		FROM(table.Users).
		WHERE(table.Users.Email.EQ(jet.String(email))).
		LIMIT(1)

	err := stmt.Query(r.db, &user)

	if errors.Is(err, qrm.ErrNoRows) {
		return entity.User{}, nil
	}

	if err != nil {
		return entity.User{}, err
	}

	return dto.EntityUserByModel(user), nil
}

func (r repository) FindUserByNetworkId(id string) (entity.User, error) {
	panic("implement me")
}

func (r repository) AttachNetwork(user entity.JwtUser, network auth.Name) {
	panic("implement me")
}

func (r repository) SaveRefreshToken(id uuid.UUID, refreshToken token.Token) {
	panic("implement me")
}

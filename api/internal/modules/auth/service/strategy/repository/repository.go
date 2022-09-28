package repository

import (
	jet "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
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
	AttachNetwork(refreshToken token.Token, network auth.Name)
	SaveRefreshToken(session model.Sessions) error
	RemoveRefreshToken(refreshToken string) error
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

func (r repository) AttachNetwork(refreshToken token.Token, network auth.Name) {
	panic("implement me")
}

func (r repository) SaveRefreshToken(session model.Sessions) error {
	stmt := table.Sessions.
		INSERT(
			table.Sessions.Token,
			table.Sessions.ExpiresAt,
			table.Sessions.UserID,
			table.Sessions.Strategy,
			table.Sessions.IP,
			table.Sessions.UserAgent,
			table.Sessions.Origin,
		).
		MODEL(session)

	_, err := stmt.Exec(r.db)

	return err
}

func (r repository) RemoveRefreshToken(refreshToken string) error {
	stmt := table.Sessions.
		DELETE().
		WHERE(table.Sessions.Token.EQ(jet.String(refreshToken)))

	_, err := stmt.Exec(r.db)

	return err
}

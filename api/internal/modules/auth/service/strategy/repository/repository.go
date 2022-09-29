package repository

import (
	"time"

	jet "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"primedivident/internal/config/consts"
	"primedivident/internal/models/app/public/model"
	"primedivident/internal/models/app/public/table"
	"primedivident/internal/modules/auth/dto"
	"primedivident/internal/modules/auth/entity"
	"primedivident/internal/modules/auth/service/strategy/auth"
	"primedivident/pkg/db/postgres"
)

type Repository interface {
	FindUserByEmail(email string) (entity.User, error)
	FindNetworkByID(networkID string, strategy auth.Name) (entity.Network, error)
	CreateUser(user model.Users) error
	SaveRefreshToken(session model.Sessions) error
	RemoveExpireRefreshToken(userID uuid.UUID) error
	RemoveLastRefreshToken(userID uuid.UUID) error
	RemoveRefreshToken(refreshToken string) error
	AttachNetwork(network model.UserNetworks) error
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

func (r repository) FindNetworkByID(networkID string, strategy auth.Name) (entity.Network, error) {
	var dest struct {
		model.UserNetworks
		model.Users
	}

	stmt := table.UserNetworks.
		SELECT(table.UserNetworks.AllColumns, table.Users.AllColumns).
		FROM(
			table.UserNetworks.
				INNER_JOIN(table.Users, table.UserNetworks.UserID.EQ(table.Users.ID)),
		).
		WHERE(jet.AND(
			table.UserNetworks.Identity.EQ(jet.String(networkID)),
			table.UserNetworks.Strategy.EQ(jet.String(strategy.String())),
		)).
		LIMIT(1)

	err := stmt.Query(r.db, &dest)

	if errors.Is(err, qrm.ErrNoRows) {
		return entity.Network{}, nil
	}

	if err != nil {
		return entity.Network{}, err
	}

	return dto.EntityUserNetworksByModel(dest.UserNetworks, dest.Users.Email, dest.Users.Name), nil
}

func (r repository) CreateUser(user model.Users) error {
	stmt := table.Users.INSERT(
		table.Users.ID,
		table.Users.Email,
		table.Users.Name,
		table.Users.Password,
		table.Users.Role,
		table.Users.Status,
	).MODEL(user)

	_, err := stmt.Exec(r.db)

	return err
}

func (r repository) SaveRefreshToken(session model.Sessions) error {
	stmt := table.Sessions.INSERT(
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
	stmt := table.Sessions.DELETE().
		WHERE(table.Sessions.Token.EQ(jet.String(refreshToken)))

	_, err := stmt.Exec(r.db)

	return err
}

func (r repository) RemoveExpireRefreshToken(userID uuid.UUID) error {
	stmt := table.Sessions.DELETE().
		WHERE(jet.AND(
			table.Sessions.UserID.EQ(jet.UUID(userID)),
			table.Sessions.ExpiresAt.LT(jet.TimestampzT(time.Now())),
		))

	_, err := stmt.Exec(r.db)

	return err
}

func (r repository) RemoveLastRefreshToken(userID uuid.UUID) error {
	stmt := table.Sessions.DELETE().
		WHERE(jet.AND(
			table.Sessions.UserID.EQ(jet.UUID(userID)),
			table.Sessions.ID.NOT_IN(
				table.Sessions.
					SELECT(table.Sessions.ID).
					WHERE(table.Sessions.UserID.EQ(jet.UUID(userID))).
					ORDER_BY(table.Sessions.CreatedAt.DESC()).
					LIMIT(consts.MaxAuthSessions),
			),
		))

	_, err := stmt.Exec(r.db)

	return err
}

func (r repository) AttachNetwork(network model.UserNetworks) error {
	stmt := table.UserNetworks.INSERT(
		table.UserNetworks.UserID,
		table.UserNetworks.Identity,
		table.UserNetworks.Strategy,
	).
		MODEL(network)

	_, err := stmt.Exec(r.db)

	return err
}

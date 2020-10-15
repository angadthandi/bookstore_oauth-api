package db

import (
	"github.com/angadthandi/bookstore_oauth-api/src/clients/cassandra"
	"github.com/angadthandi/bookstore_oauth-api/src/domain/access_token"
	"github.com/angadthandi/bookstore_oauth-api/src/utils/errors"
)

const (
	queryGetAccessToken = `SELECT access_token, user_id, client_id, expires
							FROM access_tokens WHERE access_token=?;`
	queryCreateAccessToken = `INSERT INTO access_tokens
							(access_token, user_id, client_id, expires)
							VALUES(?,?,?,?);`
	queryUpdateExpires = `UPDATE access_tokens
							SET expires = ?
							WHERE access_token=?;`
)

type DBRepository interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

type dbRepository struct {
}

func New() DBRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetByID(
	id string,
) (*access_token.AccessToken, *errors.RestErr) {
	var ret access_token.AccessToken
	err := cassandra.GetSession().
		Query(queryGetAccessToken, id).
		Scan(
			&ret.AccessToken,
			&ret.UserID,
			&ret.ClientID,
			&ret.Expires,
		)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	return &ret, nil
}

func (rr *dbRepository) Create(
	at access_token.AccessToken,
) *errors.RestErr {
	err := cassandra.GetSession().
		Query(
			queryCreateAccessToken,
			at.AccessToken,
			at.UserID,
			at.ClientID,
			at.Expires,
		).
		Exec()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

func (rr *dbRepository) UpdateExpirationTime(
	at access_token.AccessToken,
) *errors.RestErr {
	err := cassandra.GetSession().
		Query(
			queryUpdateExpires,
			at.Expires,
			at.AccessToken,
		).
		Exec()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

package db

import (
	"errors"

	"github.com/angadthandi/bookstore_oauth-api/src/clients/cassandra"
	"github.com/angadthandi/bookstore_oauth-api/src/domain/access_token"

	// "github.com/angadthandi/bookstore_oauth-api/src/utils/errors"
	"github.com/angadthandi/bookstore_utils-go/rest_errors"
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
	GetByID(string) (*access_token.AccessToken, rest_errors.RestErr)
	Create(access_token.AccessToken) rest_errors.RestErr
	UpdateExpirationTime(access_token.AccessToken) rest_errors.RestErr
}

type dbRepository struct {
}

func New() DBRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetByID(
	id string,
) (*access_token.AccessToken, rest_errors.RestErr) {
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
		// return nil, rest_errors.NewInternalServerError(err.Error(), nil)
		return nil, rest_errors.NewInternalServerError(
			"error when trying to get current id",
			errors.New("database error"),
		)
	}

	return &ret, nil
}

func (rr *dbRepository) Create(
	at access_token.AccessToken,
) rest_errors.RestErr {
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
		// return rest_errors.NewInternalServerError(err.Error(), nil)
		return rest_errors.NewInternalServerError(
			"error when trying to create new resource",
			errors.New("database error"),
		)
	}

	return nil
}

func (rr *dbRepository) UpdateExpirationTime(
	at access_token.AccessToken,
) rest_errors.RestErr {
	err := cassandra.GetSession().
		Query(
			queryUpdateExpires,
			at.Expires,
			at.AccessToken,
		).
		Exec()
	if err != nil {
		// return rest_errors.NewInternalServerError(err.Error(), nil)
		return rest_errors.NewInternalServerError(
			"error when trying to update current resource",
			errors.New("database error"),
		)
	}

	return nil
}

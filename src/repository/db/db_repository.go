package db

import (
	"github.com/angadthandi/bookstore_oauth-api/src/domain/access_token"
	"github.com/angadthandi/bookstore_oauth-api/src/utils/errors"
)

type DBRepository interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func New() DBRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetByID(
	id string,
) (*access_token.AccessToken, *errors.RestErr) {
	return nil, errors.NewInternalServerError("db conn not implemented")
}

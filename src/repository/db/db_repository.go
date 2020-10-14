package db

import (
	"fmt"

	"github.com/angadthandi/bookstore_oauth-api/src/clients/cassandra"
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
	session, err := cassandra.GetSession()
	if err != nil {
		fmt.Println("DBRepository cassandra error")
		panic(err)
	}
	defer session.Close()

	return nil, errors.NewInternalServerError("db conn not implemented")
}

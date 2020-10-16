package users

import (
	"time"

	"github.com/angadthandi/bookstore_oauth-api/src/domain/users"
	"github.com/angadthandi/bookstore_oauth-api/src/utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	usersRestClient = rest.RequestBuilder{
		// bookstore_users-api port
		BaseURL: "http://localhost:8080",
		Timeout: 100 * time.Millisecond,
	}
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type userRepository struct {
}

func New() RestUsersRepository {
	return &userRepository{}
}

func (repo *userRepository) LoginUser(
	email string,
	password string,
) (u *users.User, err *errors.RestErr) {
	return
}

package rest

import (
	"encoding/json"
	"time"

	"github.com/angadthandi/bookstore_oauth-api/src/domain/users"
	"github.com/angadthandi/bookstore_oauth-api/src/utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	usersRestClient = rest.RequestBuilder{
		// bookstore_users-api port
		BaseURL: "http://localhost:8081",
		Timeout: 100 * time.Millisecond,
	}

	InvalidRestClientErrMsg     = "invalid restclient response when trying to login"
	InvalidErrorInterfaceErrMsg = "invalid error interface when trying to login"
	UnmarshalErrMsg             = "unmarshal error when trying to login"
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
) (*users.User, *errors.RestErr) {
	req := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	resp := usersRestClient.Post("/users/login", req)
	if resp == nil || resp.Response == nil {
		return nil, errors.NewInternalServerError(
			InvalidRestClientErrMsg,
		)
	}

	if resp.StatusCode > 299 {
		var restErr errors.RestErr
		err := json.Unmarshal(resp.Bytes(), &restErr)
		if err != nil {
			return nil, errors.NewInternalServerError(
				InvalidErrorInterfaceErrMsg,
			)
		}
		return nil, &restErr
	}

	var user users.User
	err := json.Unmarshal(resp.Bytes(), &user)
	if err != nil {
		return nil, errors.NewInternalServerError(
			UnmarshalErrMsg,
		)
	}

	return &user, nil
}

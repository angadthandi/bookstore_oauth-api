package rest

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/angadthandi/bookstore_oauth-api/src/domain/users"
	// "github.com/angadthandi/bookstore_oauth-api/src/utils/errors"
	"github.com/angadthandi/bookstore_utils-go/rest_errors"
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
	LoginUser(string, string) (*users.User, rest_errors.RestErr)
}

type userRepository struct {
}

func New() RestUsersRepository {
	return &userRepository{}
}

func (repo *userRepository) LoginUser(
	email string,
	password string,
) (*users.User, rest_errors.RestErr) {
	req := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	resp := usersRestClient.Post("/users/login", req)
	if resp == nil || resp.Response == nil {
		return nil, rest_errors.NewInternalServerError(
			InvalidRestClientErrMsg,
			errors.New("restclient error"),
		)
	}

	if resp.StatusCode > 299 {
		apiErr, err := rest_errors.NewRestErrorFromBytes(resp.Bytes())
		// var restErr rest_errors.RestErr
		// err := json.Unmarshal(resp.Bytes(), &restErr)
		if err != nil {
			return nil, rest_errors.NewInternalServerError(
				InvalidErrorInterfaceErrMsg,
				err,
			)
		}
		// return nil, restErr
		return nil, apiErr
	}

	var user users.User
	err := json.Unmarshal(resp.Bytes(), &user)
	if err != nil {
		return nil, rest_errors.NewInternalServerError(
			UnmarshalErrMsg,
			errors.New("json parse error"),
		)
	}

	return &user, nil
}

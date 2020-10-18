// CHANGED
//
// this is the new code:
// https://github.com/federicoleon/bookstore_oauth-api/blob/master/src/services/access_token/access_token_service.go

package access_token

import (
	"strings"

	"github.com/angadthandi/bookstore_oauth-api/src/domain/access_token"
	"github.com/angadthandi/bookstore_oauth-api/src/repository/db"
	"github.com/angadthandi/bookstore_oauth-api/src/repository/rest"
	"github.com/angadthandi/bookstore_oauth-api/src/utils/errors"
)

// type Repository interface {
// 	GetByID(string) (*AccessToken, *errors.RestErr)
// 	Create(AccessToken) *errors.RestErr
// 	UpdateExpirationTime(AccessToken) *errors.RestErr
// }

type Service interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestErr)
	Create(
		access_token.AccessTokenRequest,
	) (*access_token.AccessToken, *errors.RestErr)
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

type service struct {
	restUsersRepo rest.RestUsersRepository
	dbRepo        db.DBRepository
}

func New(
	restUsersRepo rest.RestUsersRepository,
	dbRepo db.DBRepository,
) Service {
	return &service{
		restUsersRepo: restUsersRepo,
		dbRepo:        dbRepo,
	}
}

func (s *service) GetByID(
	accessTokenID string,
) (*access_token.AccessToken, *errors.RestErr) {
	accessTokenID = strings.TrimSpace(accessTokenID)
	if len(accessTokenID) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}

	accessToken, err := s.dbRepo.GetByID(accessTokenID)
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

func (s *service) Create(
	request access_token.AccessTokenRequest,
) (*access_token.AccessToken, *errors.RestErr) {
	err := request.Validate()
	if err != nil {
		return nil, err
	}

	// Authenticate the user against the Users API:
	user, err := s.restUsersRepo.LoginUser(
		request.Username, request.Password,
	)
	if err != nil {
		return nil, err
	}

	// Generate a new access token:
	at := access_token.GetNewAccessToken(user.ID)
	at.Generate()

	err = s.dbRepo.Create(at)
	if err != nil {
		return nil, err
	}

	return &at, nil
}

func (s *service) UpdateExpirationTime(
	at access_token.AccessToken,
) *errors.RestErr {
	err := at.Validate()
	if err != nil {
		return err
	}

	return s.dbRepo.UpdateExpirationTime(at)
}

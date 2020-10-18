package access_token

import (
	"fmt"
	"strings"
	"time"

	"github.com/angadthandi/bookstore_oauth-api/src/utils/crypto_utils"
	"github.com/angadthandi/bookstore_oauth-api/src/utils/errors"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	// for grant_type = 'password'
	Username string `json:"username"`
	Password string `json:"password"`

	// for grant_type = 'client_credentials'
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (atReq *AccessTokenRequest) Validate() *errors.RestErr {
	switch atReq.GrantType {
	case grantTypePassword:
		break
	case grantTypeClientCredentials:
		break
	default:
		return errors.NewBadRequestError("invalid grant type")
	}

	return nil
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequestError("invalid access token id")
	}

	if at.UserID <= 0 ||
		at.ClientID <= 0 ||
		at.Expires <= 0 {
		return errors.NewBadRequestError("invalid access userid or clientid or expires")
	}

	return nil
}

func GetNewAccessToken(userID int64) AccessToken {
	return AccessToken{
		UserID:  userID,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	// now := time.Now().UTC()
	// expirationTime := time.Unix(at.Expires, 0)

	// return now.After(expirationTime)
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

func (at *AccessToken) Generate() {
	at.AccessToken = crypto_utils.GetMd5(
		fmt.Sprintf("at-%d-%d-ran", at.UserID, at.Expires),
	)
}

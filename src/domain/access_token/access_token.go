package access_token

import (
	"strings"
	"time"

	"github.com/angadthandi/bookstore_oauth-api/src/utils/errors"
)

const (
	expirationTime = 24
)

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

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	// now := time.Now().UTC()
	// expirationTime := time.Unix(at.Expires, 0)

	// return now.After(expirationTime)
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

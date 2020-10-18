package rest

// import (
// 	"fmt"
// 	"net/http"
// 	"os"
// 	"testing"

// 	"github.com/mercadolibre/golang-restclient/rest"
// 	"github.com/stretchr/testify/assert"
// )

// // https://stackoverflow.com/questions/60235896/flag-provided-but-not-defined-test-v
// //
// // "github.com/mercadolibre/golang-restclient/rest"
// // library does not support the latest versions of Go.
// // It works well with GO VERSIONS <= 1.12
// //
// // so all tests will fail :(

// // TODO - maybe use a different REST Library
// // this one? - https://github.com/go-resty/resty

// func TestMain(m *testing.M) {
// 	fmt.Println("about to start test cases...")
// 	rest.StartMockupServer()
// 	os.Exit(m.Run())
// }

// func TestLoginUserTimeoutFromApi(t *testing.T) {
// 	rest.FlushMockups()
// 	rest.AddMockups(
// 		&rest.Mock{
// 			HTTPMethod:   http.MethodPost,
// 			URL:          "http://localhost:8081/users/login",
// 			ReqBody:      `{"email":"abc@test.com","password":"password"}`,
// 			RespHTTPCode: -1,
// 			RespBody:     `{}`,
// 		},
// 	)

// 	repo := userRepository{}

// 	user, err := repo.LoginUser("abc@test.com", "password")

// 	assert.Nil(t, user)
// 	assert.NotNil(t, err)
// 	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
// 	assert.EqualValues(t, InvalidRestClientErrMsg, err.Message)
// }

// func TestLoginUserInvalidErrorInterface(t *testing.T) {

// }

// func TestLoginUserInvalidLoginCredentials(t *testing.T) {

// }

// func TestLoginUserInvalidJsonResponse(t *testing.T) {

// }

// func TestLoginUserNoError(t *testing.T) {

// }

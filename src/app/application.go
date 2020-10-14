package app

import (
	"fmt"

	"github.com/angadthandi/bookstore_oauth-api/src/clients/cassandra"
	"github.com/angadthandi/bookstore_oauth-api/src/domain/access_token"
	"github.com/angadthandi/bookstore_oauth-api/src/http"
	"github.com/angadthandi/bookstore_oauth-api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session, err := cassandra.GetSession()
	if err != nil {
		fmt.Println("StartApplication cassandra error")
		panic(err)
	}
	session.Close()

	dbRepository := db.New()
	atService := access_token.New(dbRepository)
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)

	router.Run(":8080")
}

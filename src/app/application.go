package app

import (
	"github.com/angadthandi/bookstore_oauth-api/src/http"
	"github.com/angadthandi/bookstore_oauth-api/src/repository/db"
	"github.com/angadthandi/bookstore_oauth-api/src/repository/rest"
	"github.com/angadthandi/bookstore_oauth-api/src/services/access_token"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	usersRepo := rest.New()
	dbRepository := db.New()
	atService := access_token.New(usersRepo, dbRepository)
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.POST("/oauth/access_token", atHandler.Create)
	router.PUT("/oauth/access_token", atHandler.UpdateExpirationTime)

	router.Run(":8080")
}

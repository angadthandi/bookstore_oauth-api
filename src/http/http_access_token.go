package http

import (
	"net/http"

	"github.com/angadthandi/bookstore_oauth-api/src/domain/access_token"
	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetByID(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (h *accessTokenHandler) GetByID(
	c *gin.Context,
) {
	accessToken, err := h.service.GetByID(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, accessToken)
}

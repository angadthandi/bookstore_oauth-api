package http

import (
	"net/http"

	"github.com/angadthandi/bookstore_oauth-api/src/domain/access_token"
	"github.com/angadthandi/bookstore_oauth-api/src/utils/errors"
	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetByID(*gin.Context)
	Create(*gin.Context)
	UpdateExpirationTime(*gin.Context)
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

func (h *accessTokenHandler) Create(
	c *gin.Context,
) {
	var at access_token.AccessToken
	err := c.ShouldBindJSON(&at)
	if err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	createErr := h.service.Create(at)
	if createErr != nil {
		c.JSON(createErr.Status, createErr)
		return
	}

	c.JSON(http.StatusCreated, at)
}

func (h *accessTokenHandler) UpdateExpirationTime(
	c *gin.Context,
) {
	var at access_token.AccessToken
	err := c.ShouldBindJSON(&at)
	if err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	updateErr := h.service.UpdateExpirationTime(at)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}

	c.JSON(http.StatusOK, at)
}

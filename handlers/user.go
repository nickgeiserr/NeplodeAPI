package handlers

import (
	"NeplodeAPI/logger"
	"NeplodeAPI/services"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

type (
	UserHandler interface {
		GetUsers(c echo.Context) error
	}

	userHandler struct {
		services.UserService
	}
)

func (u userHandler) GetUsers(c echo.Context) error {
	r, err := services.UserService.GetUsers(u.UserService)
	claims := c.Get("claims").(*validator.ValidatedClaims)

	if err != nil {
		logger.Error("failed to get user", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, "Internal Server Issue.")
	}

	print(r)

	return c.JSON(http.StatusOK, claims.RegisteredClaims.Subject)
}

package handlers

import (
	"NeplodeAPI/logger"
	"NeplodeAPI/models"
	"NeplodeAPI/services"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

type (
	UserHandler interface {
		GetUsers(c echo.Context) error
		CreateUser(c echo.Context) error
		GetUser(c echo.Context) error
		UpdateUser(c echo.Context) error
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

func (u userHandler) GetUser(c echo.Context) error {
	uid := c.Param("uid")
	user, err := u.UserService.GetUser(uid)
	if err != nil {
		logger.Error("Something went wrong", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSON(http.StatusOK, user)
}

func (u userHandler) CreateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		logger.Error("Something went wrong with binding.", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, "Internal Server Error.")
	}

	if user.Username == "" || user.Birthday == "" {
		return c.JSON(http.StatusBadRequest, "Missing Data.")
	}

	// assign remaining values
	user.ID = c.Get("claims").(*validator.ValidatedClaims).RegisteredClaims.Subject
	user.Bio = "Hey, my name is " + user.Username + ", I'm new here!"
	user.ProfilePicture = "default_pfp_link"

	r := u.UserService.CreateUser(user)
	if r == false {
		logger.Error("Something went wrong creating the user.")
		return c.JSON(http.StatusInternalServerError, "Internal Server Issue.")
	}

	return c.JSON(http.StatusOK, user.ID)
}

func (u userHandler) UpdateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		logger.Error("Something went wrong with binding.", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, "Internal Server Error.")
	}

	if c.Get("claims").(*validator.ValidatedClaims).RegisteredClaims.Subject != user.ID {
		return c.JSON(http.StatusUnauthorized, "You can not update this user.")
	}

	r := u.UserService.UpdateUser(user)
	if r == false {
		logger.Error("Something went wrong creating the user.")
		return c.JSON(http.StatusInternalServerError, "Internal Server Issue.")
	}

	return c.JSON(http.StatusOK, user)
}

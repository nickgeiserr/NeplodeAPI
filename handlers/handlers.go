package handlers

import (
	"NeplodeAPI/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Handlers struct {
	UserHandler
}

func New(s *services.Services) *Handlers {
	return &Handlers{
		UserHandler: &userHandler{s.User},
	}
}

func SetAPI(e *echo.Echo, h *Handlers, m echo.MiddlewareFunc) {
	g := e.Group("/api")

	g.Use(m)

	// User
	g.GET("/users/:uid", h.UserHandler.GetUsers)
}

func Echo() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	return e
}

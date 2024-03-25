package handlers

import (
	"NeplodeAPI/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Handlers struct {
	UserHandler
	ChapterHandler
}

func New(s *services.Services) *Handlers {
	return &Handlers{
		UserHandler:    &userHandler{s.User},
		ChapterHandler: &chapterHandler{s.Chapter},
	}
}

func SetAPI(e *echo.Echo, h *Handlers, m echo.MiddlewareFunc) {
	g := e.Group("/api")

	g.Use(m)

	// User
	g.GET("/users/:uid", h.UserHandler.GetUser)
	g.POST("/users", h.UserHandler.CreateUser)
	g.PUT("/users", h.UserHandler.UpdateUser)

	// Chapter
	g.GET("/chapters/:chId", h.ChapterHandler.GetChapter)
	g.POST("/chapters", h.ChapterHandler.CreateChapter)
	g.PUT("/chapters/:chId", h.ChapterHandler.UpdateChapter)
	g.DELETE("/chapters/:chId", h.ChapterHandler.DeleteChapter)

	// Membership
	g.POST("/chapters/:chId/members", h.ChapterHandler.JoinChapter)
	g.DELETE("/chapters/:chId/members", h.ChapterHandler.LeaveChapter)

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

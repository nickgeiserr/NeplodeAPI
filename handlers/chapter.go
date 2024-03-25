package handlers

import (
	"NeplodeAPI/logger"
	"NeplodeAPI/models"
	"NeplodeAPI/services"
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

type (
	ChapterHandler interface {
		CreateChapter(c echo.Context) error
		GetChapter(c echo.Context) error
		JoinChapter(c echo.Context) error
		LeaveChapter(c echo.Context) error
		DeleteChapter(c echo.Context) error
		UpdateChapter(c echo.Context) error
	}

	chapterHandler struct {
		services.ChapterService
	}
)

func (ch chapterHandler) GetChapter(c echo.Context) error {
	chId := c.Param("chId")
	chapter, err := ch.ChapterService.GetChapter(chId)
	if err != nil {
		logger.Error("Something went wrong", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSON(http.StatusOK, chapter)
}

func (ch chapterHandler) CreateChapter(c echo.Context) error {
	chapter := new(models.Chapter)
	if err := c.Bind(chapter); err != nil {
		logger.Error("Something went wrong with binding.", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, "Internal Server Error.")
	}

	chapter.OwnerUserID = c.Get("claims").(*validator.ValidatedClaims).RegisteredClaims.Subject

	if !val(chapter) {
		return c.JSON(http.StatusBadRequest, "Invalid chapter data.")
	}

	uuid, err := generateUUID()
	if err != nil {
		logger.Error("failed to generate uuid", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, "Internal Server Error.")
	}

	chapter.ChapterID = uuid

	r := ch.ChapterService.CreateChapter(chapter)
	if r == false {
		logger.Error("Something went wrong creating the chapter.")
		return c.JSON(http.StatusInternalServerError, "Internal Server Issue.")
	}

	return c.JSON(http.StatusOK, chapter)
}

func val(c *models.Chapter) bool {
	defaultLogo := "default_logo_link.jpg"
	defaultBanner := "default_banner_link.jpg"

	if c.Name == "" {
		return false
	}

	if c.LogoURL == "" {
		c.LogoURL = defaultLogo
	}

	if c.BannerURL == "" {
		c.BannerURL = defaultBanner
	}

	return true
}

func generateUUID() (string, error) {
	uuid := make([]byte, 16)

	_, err := rand.Read(uuid)
	if err != nil {
		return "", err
	}

	uuid[6] = (uuid[6] & 0x0f) | 0x40
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

func (ch chapterHandler) UpdateChapter(c echo.Context) error {
	chapter := new(models.Chapter)
	if err := c.Bind(chapter); err != nil {
		logger.Error("Something went wrong with binding.", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, "Internal Server Error.")
	}

	err := ch.ChapterService.UpdateChapter(c, chapter)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return c.JSON(http.StatusNotFound, "Chapter not found.")
		}
		logger.Error("UpdateChapter failed:", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, "Internal Server Error.")
	}

	return c.JSON(http.StatusOK, "Chapter updated successfully.")
}

func (ch chapterHandler) DeleteChapter(c echo.Context) error {
	chapterID := c.Param("chapterID")

	err := ch.ChapterService.DeleteChapter(c, chapterID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Chapter not found"})
		}
		logger.Error("Error deleting chapter:", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}

	return c.JSON(http.StatusOK, "Chapter deleted successfully")
}

///
// member stuff
///

func (ch chapterHandler) JoinChapter(c echo.Context) error {
	uid := c.Get("claims").(*validator.ValidatedClaims).RegisteredClaims.Subject
	chId := c.Param("chId")

	err := ch.ChapterService.JoinChapter(uid, chId)
	if err != nil {
		// Check if the error message indicates that the user is already a member
		if err.Error() == "user is already a member of the chapter" {
			return c.JSON(http.StatusConflict, map[string]string{"error": "User is already a member of the chapter"})
		}

		// Handle other errors as internal server errors
		logger.Error("JoinChapter failed:", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}

	return c.JSON(http.StatusOK, "User successfully joined chapter.")
}

func (ch chapterHandler) LeaveChapter(c echo.Context) error {
	uid := c.Get("claims").(*validator.ValidatedClaims).RegisteredClaims.Subject
	chId := c.Param("chId")

	err := ch.ChapterService.LeaveChapter(uid, chId)
	if err != nil {
		// Check if the error is due to the user not being a member of the chapter
		if errors.Is(err, pgx.ErrNoRows) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "User is not a member of the chapter"})
		}
		// Handle other errors as internal server errors
		logger.Error("LeaveChapter failed:", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}

	return c.JSON(http.StatusOK, "User successfully left chapter.")
}

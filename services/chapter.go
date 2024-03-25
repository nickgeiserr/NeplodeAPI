package services

import (
	"NeplodeAPI/models"
	"NeplodeAPI/stores"
	"github.com/labstack/echo/v4"
)

type (
	ChapterService interface {
		CreateChapter(chapter *models.Chapter) bool
		GetChapter(chId string) (*models.Chapter, error)
		JoinChapter(uid string, chId string) error
		LeaveChapter(uid string, chId string) error
		DeleteChapter(c echo.Context, chId string) error
		UpdateChapter(c echo.Context, chapterData *models.Chapter) error
	}

	chapterService struct {
		stores *stores.ChapterStore
	}
)

func (ch chapterService) GetChapter(chId string) (*models.Chapter, error) {
	r, err := stores.ChapterStore.GetChapter(*ch.stores, chId)
	return r, err
}

func (ch chapterService) CreateChapter(chapter *models.Chapter) bool {
	r := stores.ChapterStore.CreateChapter(*ch.stores, chapter)
	return r
}

func (ch chapterService) UpdateChapter(c echo.Context, chapterData *models.Chapter) error {
	err := stores.ChapterStore.UpdateChapter(*ch.stores, c, chapterData)
	return err
}

func (ch chapterService) JoinChapter(uid string, chId string) error {
	err := stores.ChapterStore.JoinChapter(*ch.stores, uid, chId)
	return err
}

func (ch chapterService) LeaveChapter(uid string, chId string) error {
	err := stores.ChapterStore.LeaveChapter(*ch.stores, uid, chId)
	return err
}

func (ch chapterService) DeleteChapter(c echo.Context, chId string) error {
	err := stores.ChapterStore.DeleteChapter(*ch.stores, c, chId)
	return err
}

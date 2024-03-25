package stores

import (
	"NeplodeAPI/database"
	"NeplodeAPI/logger"
	"NeplodeAPI/models"
	"context"
	"errors"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"time"
)

type (
	ChapterStore interface {
		CreateChapter(chapterData *models.Chapter) bool
		GetChapter(chId string) (*models.Chapter, error)
		JoinChapter(uid string, chId string) error
		LeaveChapter(uid string, chId string) error
		DeleteChapter(c echo.Context, chId string) error
		UpdateChapter(c echo.Context, chapterData *models.Chapter) error
	}

	chapterStore struct {
	}
)

func (ch chapterStore) GetChapter(chId string) (*models.Chapter, error) {
	query := `SELECT chapter_id, name, description, owner_user_id, creation_date, modified_date, logo_url, banner_url FROM chapters WHERE chapter_id = $1`

	if database.DB == nil {
		logger.Fatal("sum went wrong")
	}
	row := database.DB.QueryRow(context.Background(), query, chId)

	chapter := &models.Chapter{}

	err := row.Scan(&chapter.ChapterID, &chapter.Name, &chapter.Description, &chapter.OwnerUserID, &chapter.CreationDate, &chapter.ModifiedDate, &chapter.LogoURL, &chapter.BannerURL)
	if err != nil {
		return nil, err
	}

	return chapter, nil
}

func (ch chapterStore) CreateChapter(chapterData *models.Chapter) bool {
	if chapterExists(chapterData.ChapterID) {
		logger.Error("Chapter already exists.")
		return false
	}

	query := "INSERT INTO chapters (chapter_id, name, description, owner_user_id, creation_date, modified_date, logo_url, banner_url) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"

	if database.DB == nil {
		logger.Fatal("Database connection is nil.")
	}

	currentTime := time.Now()

	_, err := database.DB.Query(context.Background(), query,
		chapterData.ChapterID,
		chapterData.Name,
		chapterData.Description,
		chapterData.OwnerUserID,
		currentTime,
		currentTime,
		chapterData.LogoURL,
		chapterData.BannerURL)

	if err != nil {
		logger.Error("Failed to create chapter:", zap.Error(err))
		return false
	}

	return true
}

func (ch chapterStore) DeleteChapter(c echo.Context, chId string) error {
	var exists bool
	err := database.DB.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM chapters WHERE chapter_id = $1)", chId).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return pgx.ErrNoRows // Chapter doesn't exist
	}

	ownerID, err := ch.GetChapterOwnerID(chId)
	if err != nil {
		return err
	}

	if ownerID != c.Get("claims").(*validator.ValidatedClaims).RegisteredClaims.Subject {
		return errors.New("only the owner can delete the chapter")
	}

	_, err = database.DB.Exec(context.Background(), "DELETE FROM memberships WHERE chapter_id = $1", chId)
	if err != nil {
		return err
	}

	_, err = database.DB.Exec(context.Background(), "DELETE FROM chapters WHERE chapter_id = $1", chId)
	if err != nil {
		return err
	}

	return nil
}

func (ch chapterStore) UpdateChapter(c echo.Context, chapterData *models.Chapter) error {
	var exists bool
	err := database.DB.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM chapters WHERE chapter_id = $1)", chapterData.ChapterID).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return pgx.ErrNoRows
	}

	ownerID, err := ch.GetChapterOwnerID(chapterData.ChapterID)
	if err != nil {
		return err
	}

	if ownerID != c.Get("claims").(*validator.ValidatedClaims).RegisteredClaims.Subject {
		return errors.New("only the owner can update the chapter")
	}

	_, err = database.DB.Exec(context.Background(), `
		UPDATE chapters
		SET name = $1, description = $2, logo_url = $3, banner_url = $4, modified_date = NOW()
		WHERE chapter_id = $5`,
		chapterData.Name, chapterData.Description, chapterData.LogoURL, chapterData.BannerURL, chapterData.ChapterID)
	if err != nil {
		return err
	}

	return nil
}

func (ch chapterStore) JoinChapter(uid string, chId string) error {
	query := "INSERT INTO memberships (user_id, chapter_id) VALUES ($1, $2)"
	_, err := database.DB.Exec(context.Background(), query, uid, chId)
	if err != nil {

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			// Unique violation error code (user is already a member)
			return errors.New("user is already a member of the chapter")
		}
		return err
	}

	return nil
}

func (ch chapterStore) LeaveChapter(uid string, chId string) error {
	query := "DELETE FROM memberships WHERE user_id = $1 AND chapter_id = $2"
	result, err := database.DB.Exec(context.Background(), query, uid, chId)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

func chapterExists(chapterID string) bool {
	query := "SELECT EXISTS(SELECT 1 FROM chapters WHERE chapter_id = $1)"

	// Execute the query
	var exists bool
	err := database.DB.QueryRow(context.Background(), query, chapterID).Scan(&exists)
	if err != nil {
		logger.Error("Failed to execute query:", zap.Error(err))
		return false
	}

	return exists
}

func membershipExists(uid string, chId string) bool {
	// Query to check if the membership exists
	query := "SELECT EXISTS(SELECT 1 FROM memberships WHERE user_id = $1 AND chapter_id = $2)"

	var exists bool
	err := database.DB.QueryRow(context.Background(), query, uid, chId).Scan(&exists)
	if err != nil {
		logger.Error("Failed to execute query:", zap.Error(err))
		return false
	}

	return exists
}

func (ch chapterStore) GetChapterOwnerID(chapterID string) (string, error) {
	var ownerID string
	err := database.DB.QueryRow(context.Background(), "SELECT owner_user_id FROM chapters WHERE chapter_id = $1", chapterID).Scan(&ownerID)
	if err != nil {
		return "", err
	}
	return ownerID, nil
}

package stores

import (
	"NeplodeAPI/database"
	"NeplodeAPI/logger"
	"NeplodeAPI/models"
	"context"
	"go.uber.org/zap"
)

type (
	UserStore interface {
		GetAll() ([]models.User, error)
		CreateUser(newUser *models.User) bool
		GetUser(uid string) (*models.User, error)
		UpdateUser(updatedUser *models.User) bool
	}

	userStore struct {
	}
)

func (u userStore) GetAll() ([]models.User, error) {
	var name string
	err := database.DB.QueryRow(context.Background(), "select 'Nick'").Scan(&name)
	if err != nil {
		return nil, err
	}
	return make([]models.User, 0), nil
}

func (u userStore) GetUser(uid string) (*models.User, error) {
	query := `SELECT id, username, bio, birthday, profile_picture, creation_date FROM users WHERE id = $1`

	if database.DB == nil {
		logger.Fatal("sum went wrong")
	}
	row := database.DB.QueryRow(context.Background(), query, uid)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Bio, &user.Birthday, &user.ProfilePicture, &user.CreationDate)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u userStore) CreateUser(newUser *models.User) bool {

	query := "INSERT INTO users (id, username, bio, birthday, profile_picture) VALUES ($1, $2, $3, $4, $5)"

	if database.DB == nil {
		logger.Fatal("sum went wrong")
	}

	row, err := database.DB.Query(context.Background(), query, newUser.ID, newUser.Username, newUser.Bio, newUser.Birthday, newUser.ProfilePicture)
	if err != nil {
		logger.Error("Something went wrong : ", zap.Error(err))
		return false
	}
	row.Close()
	return true
}

func (u userStore) UpdateUser(updatedUser *models.User) bool {
	query := "UPDATE users SET username = $1, bio = $2, birthday = $3, profile_picture = $4 WHERE id = $5"

	if database.DB == nil {
		logger.Fatal("Database connection is nil")
	}

	_, err := database.DB.Exec(context.Background(), query, updatedUser.Username, updatedUser.Bio, updatedUser.Birthday, updatedUser.ProfilePicture, updatedUser.ID)
	if err != nil {
		logger.Error("Failed to update user: ", zap.Error(err))
		return false
	}

	return true
}

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
		CreateUser(newUser models.User) bool
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

func (u userStore) CreateUser(newUser models.User) bool {

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

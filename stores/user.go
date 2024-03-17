package stores

import "NeplodeAPI/models"

type (
	UserStore interface {
		GetAll() ([]models.User, error)
	}

	userStore struct {
	}
)

func (u userStore) GetAll() ([]models.User, error) {
	return make([]models.User, 0), nil
}

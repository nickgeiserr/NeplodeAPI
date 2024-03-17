package services

import (
	"NeplodeAPI/models"
	"NeplodeAPI/stores"
)

type (
	UserService interface {
		GetUsers() ([]models.User, error)
	}

	userService struct {
		stores *stores.UserStore
	}
)

func (u userService) GetUsers() ([]models.User, error) {
	r, err := stores.UserStore.GetAll(*u.stores)
	return r, err
}

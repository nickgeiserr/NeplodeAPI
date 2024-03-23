package services

import (
	"NeplodeAPI/models"
	"NeplodeAPI/stores"
)

type (
	UserService interface {
		GetUsers() ([]models.User, error)
		CreateUser(user *models.User) bool
		GetUser(uid string) (*models.User, error)
		UpdateUser(updatedUser *models.User) bool
	}

	userService struct {
		stores *stores.UserStore
	}
)

func (u userService) GetUsers() ([]models.User, error) {
	r, err := stores.UserStore.GetAll(*u.stores)
	return r, err
}

func (u userService) GetUser(uid string) (*models.User, error) {
	r, err := stores.UserStore.GetUser(*u.stores, uid)
	return r, err
}

func (u userService) CreateUser(user *models.User) bool {
	r := stores.UserStore.CreateUser(*u.stores, user)
	return r
}

func (u userService) UpdateUser(updatedUser *models.User) bool {
	r := stores.UserStore.UpdateUser(*u.stores, updatedUser)
	return r
}

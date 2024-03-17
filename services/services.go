package services

import "NeplodeAPI/stores"

type Services struct {
	User UserService
}

func New(s *stores.Stores) *Services {
	return &Services{
		User: &userService{&s.User},
	}
}

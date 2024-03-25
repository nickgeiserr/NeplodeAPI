package services

import "NeplodeAPI/stores"

type Services struct {
	User    UserService
	Chapter ChapterService
}

func New(s *stores.Stores) *Services {
	return &Services{
		User:    &userService{&s.User},
		Chapter: &chapterService{&s.Chapter},
	}
}

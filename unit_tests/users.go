package main

import (
	"NeplodeAPI/logger"
	"NeplodeAPI/models"
	"NeplodeAPI/stores"
	"log"
)

func main() {
	err := logger.New()
	if err != nil {
		log.Fatal(err)
	}
	s := stores.New()

	CreateUser(s)
}

func CreateUser(s *stores.Stores) {

	newUser := models.User{
		ID:             "2",
		Username:       "testuser",
		Bio:            "Test bio",
		Birthday:       "2000-01-01",
		ProfilePicture: "profile.jpg",
		CreationDate:   "2024-03-18",
	}

	succeed := s.User.CreateUser(newUser)

	if succeed != true {
		logger.Error("IT DIDNT WORK!!")
		return
	}
	logger.Info("it worked")

}

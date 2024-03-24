package main

import (
	"NeplodeAPI/logger"
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
	logger.Info("it worked")

}

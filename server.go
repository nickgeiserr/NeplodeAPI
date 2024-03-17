package main

import (
	"NeplodeAPI/handlers"
	"NeplodeAPI/logger"
	"NeplodeAPI/middleware"
	"NeplodeAPI/services"
	"NeplodeAPI/stores"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main() {

	err := logger.New()
	if err != nil {
		log.Fatal(err)
	}

	e := handlers.Echo()

	s := stores.New()

	ss := services.New(s)
	h := handlers.New(ss)

	jwtCheck, err := middleware.JwtMiddleware()
	if err != nil {
		logger.Fatal("failed to set JWT middleware", zap.Error(err))
	}

	handlers.SetAPI(e, h, jwtCheck)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "1515"
	}

	logger.Fatal("failed to start server", zap.Error(e.Start(":"+PORT)))
}

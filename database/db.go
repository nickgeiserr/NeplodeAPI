package database

import (
	"NeplodeAPI/logger"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
)

var DB *pgxpool.Pool

func init() {

	if err := godotenv.Load(); err != nil {
		logger.Error("Oops! Sum went wrong", zap.Error(err))
	}

	pool, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		logger.Error("Failed to create Database Pool.")
		return
	}

	DB = pool

}

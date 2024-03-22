package database

import (
	"NeplodeAPI/logger"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

var DB *pgxpool.Pool

func init() {
	pool, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		logger.Error("Failed to create Database Pool.")
		return
	}

	DB = pool

}

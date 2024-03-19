package database

import (
	"NeplodeAPI/logger"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func init() {
	pool, err := pgxpool.New(context.Background(), "postgres://neplode_admin:c8gxiTEjCq3ppMahKQn4Q5o7IUouqTRi@dpg-cnn4s9gl6cac73ferfrg-a.ohio-postgres.render.com/neplode_db")
	if err != nil {
		logger.Error("Failed to create Database Pool.")
		return
	}

	DB = pool

}

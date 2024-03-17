package database

import (
	"context"
	"github.com/jackc/pgx/v4"
	"os"
)

var DB *pgx.Conn

func init() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DBURL"))
	if err != nil {
		panic(err)
	}
	DB = conn
}

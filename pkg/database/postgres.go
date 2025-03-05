package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresDatabase struct {
	Pool *pgxpool.Pool
}

func New(dbUrl string) *PostgresDatabase {
	pool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		panic("Failed to connect database")
	}
	return &PostgresDatabase{Pool: pool}
}

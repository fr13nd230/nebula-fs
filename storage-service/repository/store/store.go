package store

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDB(driver string) (*Queries, error) {
	pgpool, err := pgxpool.New(context.Background(), driver)
	if err != nil {
		return nil, err
	}
	conn, err := pgpool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	return New(conn), nil
}
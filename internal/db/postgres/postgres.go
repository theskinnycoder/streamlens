package postgres

import (
	"context"
	"errors"
	"os"

	"github.com/jackc/pgx/v5"
)

func NewConnection(ctx context.Context) (*pgx.Conn, error) {
	dbUrl := os.Getenv("POSTGRES_URL")
	if dbUrl == "" {
		return nil, errors.New("POSTGRES_URL is not set")
	}
	return pgx.Connect(ctx, dbUrl)
}

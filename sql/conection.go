package sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateConnect(ctx context.Context) (*pgx.Conn, error) {
	return pgx.Connect(ctx, "postgres://smon:2608@localhost:5432/postgres")
}

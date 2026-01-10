package create_table

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(conn pgx.Conn, ctx context.Context) error {
	queryStr := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR NOT NULL,
		age INTEGER NOT NULL
	);
	`
	_, err := conn.Exec(ctx, queryStr)
	return err
}

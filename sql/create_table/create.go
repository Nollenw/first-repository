package create_table

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(conn pgx.Conn, ctx context.Context) error {
	queryStr := `
	CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		title VARCHAR (100) NOT NULL,
		description VARCHAR(200) NOT NULL,
		completed BOOLEAN NOT NULL,
		time_created TIMESTAMP NOT NULL,
		time_completed TIMESTAMP

	);
	`
	_, err := conn.Exec(ctx, queryStr)
	return err
}

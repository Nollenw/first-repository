package create_user

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateUser(conn pgx.Conn, ctx context.Context) error {
	queryStr := `
	INSERT INTO users (name, age)
	VALUES ('Nollen', 25)
	`
	_, err := conn.Exec(ctx, queryStr)
	return err
}

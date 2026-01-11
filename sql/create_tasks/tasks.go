package create_user

import (
	"context"
	"study/handlers"

	"github.com/jackc/pgx/v5"
)

func CreateTask(conn pgx.Conn, ctx context.Context, h *handlers.Task) error {
	queryStr := `
	INSERT INTO tasks (title, description, completed, time_created)
	VALUES ($1,$2,$3,$4);
	`
	_, err := conn.Exec(ctx, queryStr, h.Title, h.Description, h.IsCompleted, h.TimeCreate)
	return err
}

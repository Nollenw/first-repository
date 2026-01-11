package create_user

import (
	"context"
	maketask "study/makeTask"

	"github.com/jackc/pgx/v5"
)

func CreateTask(conn pgx.Conn, ctx context.Context, t *maketask.Task) error {
	queryStr := `
	INSERT INTO tasks (title, description, completed, time_created)
	VALUES ($1,$2,$3,$4);
	`
	_, err := conn.Exec(ctx, queryStr,
		t.Title,
		t.Description,
		t.IsCompleted,
		t.TimeCreate)
	return err
}

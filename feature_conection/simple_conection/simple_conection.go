package simple_conection

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func CheckConection() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://postgres:2608@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}
	if err := conn.Ping(ctx); err != nil {
		panic(err)
	}
	fmt.Println("Conect!")
}

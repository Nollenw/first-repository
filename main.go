package main

import (
	"context"
	"fmt"
	"study/sql/create_table"
	"study/sql/create_user"
	"study/sql/simple_conection"
)

func main() {
	ctx := context.Background()
	conn, err := simple_conection.CreateConnect(ctx)
	if err != nil {
		panic(err)
	}
	if err := create_table.CreateTable(*conn, ctx); err != nil {
		panic(err)
	}
	if err := create_user.CreateUser(*conn, ctx); err != nil {
		panic(err)
	}
	fmt.Println("successfully!")
}

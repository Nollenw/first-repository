package main

import (
	"context"
	"fmt"
	"net/http"
	"study/handlers"
	"study/sql"
)

func main() {
	ctx := context.Background()
	conn, err := sql.CreateConnect(ctx)
	if err != nil {
		panic(err)
	}
	if err := sql.CreateTable(*conn, ctx); err != nil {
		panic(err)
	}
	http.HandleFunc("/task", handlers.CreateTaskHandler(conn))
	fmt.Println("successfully!")
	if err := http.ListenAndServe(":9091", nil); err != nil {
		panic(err)
	}
}

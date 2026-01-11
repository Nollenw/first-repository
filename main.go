package main

import (
	"context"
	"fmt"
	"net/http"
	"study/handlers"
	simple_conection "study/sql/conection"
	"study/sql/create_table"
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
	http.HandleFunc("/task", handlers.CreateTaskHandler(conn))
	fmt.Println("successfully!")
	if err := http.ListenAndServe(":9091", nil); err != nil {
		panic(err)
	}
}

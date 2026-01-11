package main

import (
	"context"
	"fmt"
	"net/http"
	"study/handlers"
	simple_conection "study/sql/conection"
	"study/sql/create_table"
	create_user "study/sql/create_tasks"
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
	task := handlers.NewTask()
	if err := create_user.CreateTask(*conn, ctx, task); err != nil {
		panic(err)
	}
	http.HandleFunc("/task", handlers.AcceptinHttp)
	fmt.Println("successfully!")
	if err := http.ListenAndServe(":9091", nil); err != nil {
		panic(err)
	}
}

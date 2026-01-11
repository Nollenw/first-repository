package handlers

import (
	"encoding/json"
	"net/http"
	maketask "study/makeTask"
	create_user "study/sql/create_tasks"
	"time"

	"github.com/jackc/pgx/v5"
)

func CreateTaskHandler(conn *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var task maketask.Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		task.TimeCreate = time.Now()
		task.IsCompleted = false

		if err := create_user.CreateTask(*conn, r.Context(), &task); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(task)
	}
}

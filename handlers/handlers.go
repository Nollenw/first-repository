package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

type Task struct {
	Title       string
	Description string
	IsCompleted bool
	TimeCreate  time.Time
	TimeComplet *time.Time
}

func NewTask(title string, description string) *Task {
	return &Task{
		Title:       title,
		Description: description,
		IsCompleted: false,
		TimeCreate:  time.Now(),
	}
}

func AcceptinHttp(w http.ResponseWriter, r *http.Request) {
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	task.TimeCreate = time.Now()

	json.NewEncoder(w).Encode(task)
}

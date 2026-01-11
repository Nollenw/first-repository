package maketask

import "time"

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

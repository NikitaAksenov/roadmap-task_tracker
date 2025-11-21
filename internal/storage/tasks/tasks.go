package tasks

import (
	"errors"
	"time"
)

type Task struct {
	ID          int
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

const (
	StatusToDo       = "todo"
	StatusInProgress = "in-progress"
	StatusDone       = "done"
)

type TasksStorage interface {
	Add(description string) (int, error)
	Update(id int, description string) error
	Delete(id int) error
	GetAll(status *string) ([]Task, error)
}

var (
	ErrTaskNotFound   = errors.New("task not found")
	ErrTaskNotUpdated = errors.New("task not updated")
)

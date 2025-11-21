package tasks

import "time"

type Task struct {
	ID          int
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TasksStorage interface {
	Add(description string) error
	Update(id int, description string) error
	Delete(id int) error
	List(status *string) error
}

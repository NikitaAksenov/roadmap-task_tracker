package tasks

import (
	"errors"
	"strings"
	"time"
)

type Task struct {
	ID          int
	Description string
	Status      TaskStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TaskStatus int

const (
	StatusToDo TaskStatus = iota
	StatusInProgress
	StatusDone
)

const (
	statusStringToDo       = "todo"
	statusStringInProgress = "in-progress"
	statusStringDone       = "done"
)

var taskStatusString = map[TaskStatus]string{
	StatusToDo:       statusStringToDo,
	StatusInProgress: statusStringInProgress,
	StatusDone:       statusStringDone,
}

func (ts TaskStatus) String() string {
	return taskStatusString[ts]
}

func TaskStatusFromString(s string) (TaskStatus, error) {
	switch strings.TrimSpace(strings.ToLower(s)) {
	case statusStringToDo:
		return StatusToDo, nil
	case statusStringInProgress:
		return StatusInProgress, nil
	case statusStringDone:
		return StatusDone, nil
	default:
		return 0, ErrInvalidStatus
	}
}

type TasksStorage interface {
	Add(description string) (int, error)
	Update(id int, description string) error
	Delete(id int) error
	GetAll(status *TaskStatus) ([]Task, error)
	Mark(id int, status TaskStatus) error
}

var (
	ErrTaskNotFound   = errors.New("task not found")
	ErrTaskNotUpdated = errors.New("task not updated")
	ErrSameStatus     = errors.New("task is already in that status")
	ErrInvalidStatus  = errors.New("invalid task status")
)

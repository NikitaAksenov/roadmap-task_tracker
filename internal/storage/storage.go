package storage

import "github.com/NikitaAksenov/roadmap-task_tracker/internal/storage/tasks"

type Storage struct {
	Tasks tasks.TasksStorage
}

func NewStorage() Storage {
	return Storage{
		Tasks: tasks.NewStorageJSON("tasks.json"),
	}
}

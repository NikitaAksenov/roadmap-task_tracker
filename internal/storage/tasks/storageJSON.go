package tasks

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type TasksStorageJSON struct {
	filepath string
}

func NewStorageJSON(filepath string) *TasksStorageJSON {
	return &TasksStorageJSON{
		filepath: filepath,
	}
}

type taskJSON struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (tj taskJSON) toTask() (Task, error) {
	status, err := TaskStatusFromString(tj.Status)
	if err != nil {
		return Task{}, err
	}

	return Task{
		ID:          tj.ID,
		Description: tj.Description,
		Status:      status,
		CreatedAt:   tj.CreatedAt,
		UpdatedAt:   tj.UpdatedAt,
	}, nil
}

func (t Task) toTaskJSON() taskJSON {
	return taskJSON{
		ID:          t.ID,
		Description: t.Description,
		Status:      t.Status.String(),
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}

func (ts *TasksStorageJSON) readJSON() ([]Task, error) {
	fileBytes, err := os.ReadFile(ts.filepath)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	if len(fileBytes) == 0 {
		return make([]Task, 0), nil
	}

	tasksJSON := make([]taskJSON, 0)
	err = json.Unmarshal(fileBytes, &tasksJSON)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	tasks := make([]Task, len(tasksJSON))
	for i, taskJSON := range tasksJSON {
		tasks[i], err = taskJSON.toTask()
		if err != nil {
			return nil, fmt.Errorf("error converting JSON to task: %v", err)
		}
	}

	return tasks, nil
}

func (ts *TasksStorageJSON) writeJSON(tasks []Task) error {
	tasksJSON := make([]taskJSON, len(tasks))
	for i, task := range tasks {
		tasksJSON[i] = task.toTaskJSON()
	}

	jsonData, err := json.MarshalIndent(tasksJSON, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling JSON: %v", err)
	}

	file, err := os.OpenFile(ts.filepath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("error opening/creating file: %v", err)
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return nil
}

func (ts *TasksStorageJSON) Add(description string) (int, error) {
	tasks, err := ts.readJSON()
	if err != nil {
		return 0, fmt.Errorf("error during Add: %v", err)
	}

	lastTaskID := 0
	if len(tasks) > 0 {
		lastTaskID = tasks[len(tasks)-1].ID
	}

	task := Task{
		ID:          lastTaskID + 1,
		Description: description,
		Status:      StatusToDo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, task)

	err = ts.writeJSON(tasks)
	if err != nil {
		return 0, fmt.Errorf("error during Add: %v", err)
	}

	return task.ID, nil
}

func (ts *TasksStorageJSON) Update(id int, description string) error {
	tasks, err := ts.readJSON()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Description = description

			if tasks[i] == task {
				return ErrTaskNotUpdated
			}

			err = ts.writeJSON(tasks)
			if err != nil {
				return err
			}

			return nil
		}
	}

	return ErrTaskNotFound
}

func (ts *TasksStorageJSON) Delete(id int) error {
	tasks, err := ts.readJSON()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)

			err = ts.writeJSON(tasks)
			if err != nil {
				return err
			}

			return nil
		}
	}

	return ErrTaskNotFound
}

func (ts *TasksStorageJSON) GetAll(status *TaskStatus) ([]Task, error) {
	tasks, err := ts.readJSON()
	if err != nil {
		return nil, err
	}

	if status == nil {
		return tasks, nil
	}

	filteredTasks := make([]Task, 0, len(tasks))

	for _, task := range tasks {
		if status != nil && task.Status != *status {
			continue
		}

		filteredTasks = append(filteredTasks, task)
	}

	return filteredTasks, nil
}

func (ts *TasksStorageJSON) Mark(id int, status TaskStatus) error {
	tasks, err := ts.readJSON()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == id {
			if task.Status == status {
				return ErrSameStatus
			}

			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()

			err = ts.writeJSON(tasks)
			if err != nil {
				return err
			}

			return nil
		}
	}

	return ErrTaskNotFound
}

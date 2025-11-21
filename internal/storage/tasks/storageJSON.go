package tasks

type TasksStorageJSON struct {
	filepath string
}

func NewStorageJSON(filepath string) *TasksStorageJSON {
	return &TasksStorageJSON{
		filepath: filepath,
	}
}

func (ts *TasksStorageJSON) Add(description string) error {
	return nil
}

func (ts *TasksStorageJSON) Update(id int, description string) error {
	return nil
}

func (ts *TasksStorageJSON) Delete(id int) error {
	return nil
}

func (ts *TasksStorageJSON) List(status *string) error {
	return nil
}

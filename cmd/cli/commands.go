package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/NikitaAksenov/roadmap-task_tracker/internal/storage/tasks"
	"github.com/NikitaAksenov/roadmap-task_tracker/internal/validator"
)

func (app *application) commandAdd(args []string) {
	flagSet := flag.NewFlagSet("add", flag.ExitOnError)
	parameterDescription := flagSet.String("description", "", "Description for new task")
	flagSet.Parse(args)

	v := validator.New()

	ValidateParameterDescription(v, flagSet, parameterDescription)

	if !v.Valid() {
		fmt.Println(v.PrettyString())
		return
	}

	description := *parameterDescription

	id, err := app.Storage.Tasks.Add(description)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Task added successfully (ID: %d)", id)
}

func (app *application) commandUpdate(args []string) {
	flagSet := flag.NewFlagSet("update", flag.ExitOnError)
	parameterID := flagSet.Int("id", 0, "ID of task to update")
	parameterDescription := flagSet.String("description", "", "New tasks description")
	flagSet.Parse(args)

	v := validator.New()

	ValidateParameterID(v, flagSet, parameterID)
	ValidateParameterDescription(v, flagSet, parameterDescription)

	if !v.Valid() {
		fmt.Println(v.PrettyString())
		return
	}

	id := *parameterID
	description := *parameterDescription

	err := app.Storage.Tasks.Update(id, description)
	if err != nil {
		switch {
		case errors.Is(err, tasks.ErrTaskNotFound):
			fmt.Printf("Failed to find task with ID: %d", id)
		case errors.Is(err, tasks.ErrTaskNotUpdated):
			fmt.Printf("Passed parameters are the same for the task with ID: %d", id)
		default:
			fmt.Printf("Error during update: %v", err)
		}
		return
	}

	fmt.Printf("Task #%d updated successfully", id)
}

func (app *application) commandDelete(args []string) {
	flagSet := flag.NewFlagSet("update", flag.ExitOnError)
	parameterID := flagSet.Int("id", 0, "ID of task to delete")
	flagSet.Parse(args)

	v := validator.New()

	ValidateParameterID(v, flagSet, parameterID)

	if !v.Valid() {
		fmt.Println(v.PrettyString())
		return
	}

	id := *parameterID

	err := app.Storage.Tasks.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, tasks.ErrTaskNotFound):
			fmt.Printf("Failed to find task with ID: %d", id)
		default:
			fmt.Printf("Error during delete: %v", err)
		}
		return
	}

	fmt.Printf("Task #%d deleted successfully", id)
}

func (app *application) commandList(args []string) {
	flagSet := flag.NewFlagSet("list", flag.ExitOnError)
	parameterStatus := flagSet.String("status", "", "Status for task to be filtered on (optional)")
	flagSet.Parse(args)

	var status *tasks.TaskStatus
	if IsFlagPassedInSet(flagSet, "status") {
		statusValue, err := tasks.TaskStatusFromString(*parameterStatus)
		if err != nil {
			switch {
			case errors.Is(err, tasks.ErrInvalidStatus):
				fmt.Printf("Status %s is invalid\n", *parameterStatus)
			default:
				fmt.Println(err)
			}
			return
		}

		status = &statusValue
	}

	tasks, err := app.Storage.Tasks.GetAll(status)
	if err != nil {
		fmt.Printf("Error during list: %v", err)
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	fmt.Printf("%-3s %-25s %-11s %10s %10s\n", "ID", "Description", "Status", "Created At", "Updated At")
	for _, task := range tasks {
		fmt.Printf("%-3d %-25s %-11s %10s %10s\n", task.ID, task.Description, task.Status, task.CreatedAt.Format("2006-01-02"), task.UpdatedAt.Format("2006-01-02"))
	}
}

func (app *application) commandMark(args []string) {
	flagSet := flag.NewFlagSet("mark", flag.ExitOnError)
	parameterID := flagSet.Int("id", 0, "ID of task to set new status")
	parameterStatus := flagSet.String("status", "", "New status for task")
	flagSet.Parse(args)

	v := validator.New()

	ValidateParameterID(v, flagSet, parameterID)
	ValidateParameterStatus(v, flagSet, parameterStatus)

	if !v.Valid() {
		fmt.Println(v.PrettyString())
		return
	}

	id := *parameterID

	status, err := tasks.TaskStatusFromString(*parameterStatus)
	if err != nil {
		switch {
		case errors.Is(err, tasks.ErrInvalidStatus):
			fmt.Printf("Status %s is invalid\n", *parameterStatus)
		default:
			fmt.Println(err)
		}
		return
	}

	err = app.Storage.Tasks.Mark(id, status)
	if err != nil {
		fmt.Printf("Error during mark: %v", err)
		return
	}

	fmt.Printf("Successfully marked task #%d to status %s", id, status)
}

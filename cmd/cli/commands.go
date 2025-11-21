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
	fmt.Println("delete")
}

func (app *application) commandList(args []string) {
	fmt.Println("list")
}

package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/NikitaAksenov/roadmap-task_tracker/internal/storage/tasks"
)

func (app *application) commandAdd(args []string) {
	flagSet := flag.NewFlagSet("add", flag.ExitOnError)
	flagDescription := flagSet.String("description", "", "Description for new task")
	flagSet.Parse(args)

	if !IsFlagPassedInSet(flagSet, "description") {
		fmt.Println("Parameter [-description] must be provided")
		return
	}

	id, err := app.Storage.Tasks.Add(*flagDescription)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Task added successfully (ID: %d)", id)
}

func (app *application) commandUpdate(args []string) {
	flagSet := flag.NewFlagSet("update", flag.ExitOnError)
	flagID := flagSet.Int("id", 0, "ID of task to update")
	flagDescription := flagSet.String("description", "", "New tasks description")
	flagSet.Parse(args)

	if !IsFlagPassedInSet(flagSet, "id") {
		fmt.Println("Parameter [-id] must be provided")
		return
	}

	if !IsFlagPassedInSet(flagSet, "description") {
		fmt.Println("Parameter [-description] must be provided")
		return
	}

	id := *flagID
	description := *flagDescription

	if id <= 0 {
		fmt.Println("Parameter [-id] must be > 0")
		return
	}

	if description == "" {
		fmt.Println("Parameter [-description] must not be empty")
		return
	}

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

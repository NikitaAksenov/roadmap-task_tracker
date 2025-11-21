package main

import (
	"flag"
	"fmt"
)

func (app *application) commandAdd(args []string) {
	flagSet := flag.NewFlagSet("add", flag.ExitOnError)
	flagDescription := flagSet.String("description", "", "New task description")
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
	fmt.Println("update")
}

func (app *application) commandDelete(args []string) {
	fmt.Println("delete")
}

func (app *application) commandList(args []string) {
	fmt.Println("list")
}

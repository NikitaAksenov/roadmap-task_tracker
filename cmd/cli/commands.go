package main

import "fmt"

func (app *application) commandAdd(args []string) {
	fmt.Println("add")

	err := app.Storage.Tasks.Add("test")
	if err != nil {
		fmt.Println(err)
		return
	}
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

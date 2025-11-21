package main

import "fmt"

func (app *application) commandAdd(args []string) {
	fmt.Println("add")
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

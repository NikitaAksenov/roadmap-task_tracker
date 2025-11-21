package main

import (
	"fmt"
	"os"

	"github.com/NikitaAksenov/roadmap-task_tracker/internal/storage"
)

type application struct {
	Storage storage.Storage
}

func main() {
	app := &application{
		Storage: storage.NewStorage(),
	}

	if len(os.Args) < 2 {
		fmt.Println("No command passed")
		return
	}

	command := os.Args[1]
	args := os.Args[2:]

	router := make(map[string]func([]string))
	router["add"] = app.commandAdd
	router["update"] = app.commandUpdate
	router["delete"] = app.commandDelete
	router["list"] = app.commandList

	if commandFunc, ok := router[command]; ok {
		commandFunc(args)
	} else {
		fmt.Printf("Unknown command [%s]\n", command)
	}
}

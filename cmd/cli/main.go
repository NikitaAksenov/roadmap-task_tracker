package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/NikitaAksenov/roadmap-task_tracker/internal/router"
	"github.com/NikitaAksenov/roadmap-task_tracker/internal/storage"
)

type application struct {
	Router  router.Router
	Storage storage.Storage
}

func main() {
	app := &application{
		Router:  *router.NewRouter(),
		Storage: storage.NewStorage(),
	}

	app.prepareCommands()

	if len(os.Args) < 2 {
		fmt.Println("No command passed")
		app.PrintAllowedCommands()
		return
	}

	command := os.Args[1]
	args := os.Args[2:]

	err := app.Router.Execute(command, args)
	if err != nil {
		switch {
		case errors.Is(err, router.ErrCommandNotExists):
			fmt.Printf("Unknown command [%s]\n", command)
			app.PrintAllowedCommands()
		default:
			fmt.Println("Error:", err)
		}
	}
}

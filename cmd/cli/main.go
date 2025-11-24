package main

import (
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

	if len(os.Args) < 2 {
		fmt.Println("No command passed")
		return
	}

	command := os.Args[1]
	args := os.Args[2:]

	app.prepareCommands()

	app.Router.Execute(command, args)
}

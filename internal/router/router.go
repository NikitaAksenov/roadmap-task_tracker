package router

import "errors"

var (
	ErrCommandAlreadyExists = errors.New("command already exists")
	ErrCommandNotExists     = errors.New("command not exist")
)

type Router struct {
	commands map[string]func([]string)
}

func NewRouter() *Router {
	r := &Router{
		commands: make(map[string]func([]string)),
	}

	return r
}

func (r *Router) AddCommand(commandName string, command func([]string)) error {
	if _, ok := r.commands[commandName]; ok {
		return ErrCommandAlreadyExists
	}

	r.commands[commandName] = command

	return nil
}

func (r *Router) Execute(commandName string, args []string) error {
	if command, ok := r.commands[commandName]; ok {
		command(args)

		return nil
	}

	return ErrCommandNotExists
}

func (r *Router) GetAllowedCommands() []string {
	allowedCommands := make([]string, 0, len(r.commands))

	for k := range r.commands {
		allowedCommands = append(allowedCommands, k)
	}

	return allowedCommands
}

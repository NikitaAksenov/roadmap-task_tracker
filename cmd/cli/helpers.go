package main

import (
	"flag"
	"fmt"
)

func IsFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func IsFlagPassedInSet(set *flag.FlagSet, name string) bool {
	found := false
	set.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func (app *application) PrintAllowedCommands() {
	fmt.Println("Allowed commands:")
	for _, command := range app.Router.GetAllowedCommands() {
		fmt.Println("-", command)
	}
}

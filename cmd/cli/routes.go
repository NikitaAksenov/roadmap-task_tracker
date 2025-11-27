package main

func (app *application) prepareCommands() {
	app.Router.AddCommand("add", app.commandAdd)
	app.Router.AddCommand("delete", app.commandDelete)
	app.Router.AddCommand("list", app.commandList)
	app.Router.AddCommand("mark", app.commandMark)
	app.Router.AddCommand("update", app.commandUpdate)
}

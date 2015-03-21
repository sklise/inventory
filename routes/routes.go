package routes

import "github.com/sklise/inventory/config"

var App *config.App

func Setup(app *config.App) {
	App = app
}

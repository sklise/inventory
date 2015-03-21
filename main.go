package main

import (
	"net/http"
	"os"

	"github.com/sklise/inventory/config"
	"github.com/sklise/inventory/routes"
)

var App *config.App

func main() {
	App = config.NewApp()
	routes.Setup(App)

	App.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		App.Render.HTML(w, 200, "index", nil)
	}).Methods("Get")

	// Things
	App.Router.HandleFunc("/things", routes.ThingsIndex).Methods("Get")
	App.Router.HandleFunc("/things/new", routes.ThingsNew).Methods("Get")
	App.Router.HandleFunc("/things", routes.ThingsCreate).Methods("Post")
	App.Router.HandleFunc("/things/{id}", routes.ThingsShow).Methods("Get")
	App.Router.HandleFunc("/things/{id}", routes.ThingsUpdate).Methods("Put")
	App.Router.HandleFunc("/things/{id}", routes.ThingsDestroy).Methods("Delete")

	// Authors
	App.Router.HandleFunc("/authors", routes.AuthorsIndex).Methods("Get")
	App.Router.HandleFunc("/authors/new", routes.AuthorsNew).Methods("Get")
	App.Router.HandleFunc("/authors", routes.AuthorsCreate).Methods("Post")
	App.Router.HandleFunc("/authors/{id}", routes.AuthorsShow).Methods("Get")

	// Launch Negroni with all middleware and run on this port.
	App.Negroni.Run(":" + os.Getenv("PORT"))
}

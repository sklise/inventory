package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/sklise/inventory/config"
	"github.com/sklise/inventory/models"
)

var App *config.App

func Setup(app *config.App) {
	App = app
}

func ThingsIndex(w http.ResponseWriter, r *http.Request) {
	things := []models.Thing{}

	App.DB.Find(&things)
	App.Render.HTML(w, 200, "things/index", things)
}

func ThingsNew(w http.ResponseWriter, r *http.Request) {
	authors := []models.Author{}
	App.DB.Find(&authors)
	App.Render.HTML(w, 200, "things/new", authors)
}

func ThingsCreate(w http.ResponseWriter, r *http.Request) {
	// Parse form values
	err1 := r.ParseForm()
	if err1 != nil {
		fmt.Println("Cannot parse form")
	}

	title := r.FormValue("title")
	year, _ := strconv.ParseInt(r.FormValue("year"), 10, 64)
	author_id, _ := strconv.ParseInt(r.FormValue("author_id"), 10, 64)

	thing := models.Thing{
		Title:    title,
		Year:     year,
		AuthorId: author_id,
	}

	err := App.DB.Create(&thing)
	if err != nil {
		App.Render.HTML(w, 500, "error", err)
		return
	}

	things := []models.Thing{}
	App.DB.Find(&things)
	App.Render.HTML(w, 200, "things/index", things)
}

func ThingsShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	thing := models.Thing{}
	App.DB.Where("id = ?", vars["id"]).First(&thing)

	if thing.Id == 0 {
		App.Render.HTML(w, 404, "error", "404 Could not find requested thing")
	} else {
		App.Render.HTML(w, 200, "things/show", thing)
	}
}

func ThingsUpdate(w http.ResponseWriter, r *http.Request) {
	vars := context.Get(r, "params")
	fmt.Fprintf(w, "%s", vars)
}

func ThingsDestroy(w http.ResponseWriter, r *http.Request) {
	vars := context.Get(r, "params")
	fmt.Fprintf(w, "ID: %v", vars)
}

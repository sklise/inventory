package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sklise/inventory/models"
)

// Index for Authors
func AuthorsIndex(w http.ResponseWriter, r *http.Request) {
	authors := []models.Author{}
	App.DB.Order("LOWER(name) asc").Find(&authors)

	App.Render.HTML(w, 200, "authors/index", authors)
}

// New Author Form
func AuthorsNew(w http.ResponseWriter, r *http.Request) {
	App.Render.HTML(w, 200, "authors/new", "")
}

// Create Author
func AuthorsCreate(w http.ResponseWriter, r *http.Request) {
	err1 := r.ParseForm()
	if err1 != nil {
		App.Render.HTML(w, 500, "error", 500)
		return
	}

	author := models.Author{
		Name: r.FormValue("name"),
	}

	App.DB.Create(&author)

	http.Redirect(w, r, "/authors", http.StatusFound)

}

// Show Author
func AuthorsShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	author := models.Author{}
	things := []models.Thing{}

	if App.DB.Where("id = ?", vars["id"]).First(&author).RecordNotFound() {
		App.Render.HTML(w, 404, "error", "404 Could not find requested author")
		return
	}

	App.DB.Model(&author).Order("year desc").Related(&things)

	data := models.AuthorAndThings{
		Author: author,
		Things: things,
	}

	App.Render.HTML(w, 200, "authors/show", data)
}

func AuthorsUpdate(w http.ResponseWriter, r *http.Request)  {}
func AuthorsDestroy(w http.ResponseWriter, r *http.Request) {}

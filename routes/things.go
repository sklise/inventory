package routes

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sklise/inventory/models"
)

func ThingsIndex(w http.ResponseWriter, r *http.Request) {
	things := []models.Thing{}
	App.DB.Find(&things)
	App.Render.HTML(w, http.StatusOK, "things/index", things)
}

func ThingsNew(w http.ResponseWriter, r *http.Request) {
	authors := []models.Author{}
	App.DB.Find(&authors)
	App.Render.HTML(w, http.StatusOK, "things/new", authors)
}

func ThingsCreate(w http.ResponseWriter, r *http.Request) {
	// Parse form values
	err1 := r.ParseForm()
	if err1 != nil {
		App.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": "Cannot parse form"})
		return
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
		App.Render.JSON(w, http.StatusInternalServerError, err)
		return
	}

	App.Render.JSON(w, http.StatusOK, thing)
}

func ThingsShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	thing := models.Thing{}
	App.DB.Where("id = ?", params["id"]).First(&thing)

	if thing.Id == 0 {
		App.Render.HTML(w, 404, "error", "404 Could not find requested thing")
	} else {
		App.Render.HTML(w, http.StatusOK, "things/show", thing)
	}
}

func ThingsUpdate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// Parse form values
	err1 := r.ParseForm()
	if err1 != nil {
		App.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": "Cannot parse form"})
		return
	}

	thing := models.Thing{}
	App.DB.Where("id = ?", params["id"]).First(&thing)

	if thing.Id == 0 {
		App.Render.JSON(w, 404, map[string]string{"error": "404 Could not find requested thing"})
		return
	} else {
		title := r.FormValue("title")
		year, _ := strconv.ParseInt(r.FormValue("year"), 10, 64)
		author_id, _ := strconv.ParseInt(r.FormValue("author_id"), 10, 64)

		thing.Title = title
		thing.Year = year
		thing.AuthorId = author_id
		App.DB.Save(&thing)

		App.Render.JSON(w, http.StatusOK, thing)
		return
	}
}

func ThingsDestroy(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	thing := models.Thing{}

	App.DB.Where("id = ?", params["id"]).First(&thing)

	if thing.Id == 0 {
		App.Render.JSON(w, http.StatusNotFound, map[string]string{"error": "content not found"})
		return
	}

	App.DB.Delete(thing)

	App.Render.JSON(w, http.StatusOK, thing)
}

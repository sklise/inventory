package config

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/unrolled/render"
)

type App struct {
	Negroni *negroni.Negroni
	Router  *mux.Router
	Render  *render.Render
	DB      *gorm.DB
}

func newDB() *gorm.DB {
	db, err := gorm.Open("postgres", "dbname=inventory sslmode=disable")

	// Uncomment this to enable DB loggin
	//db.LogMode(true)

	if err != nil {
		fmt.Printf("DB connection error: %v\n", err)
	}
	return &db
}

func NewApp() *App {
	db := newDB()

	// Use negroni for middleware
	ne := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
		negroni.NewStatic(http.Dir("public")),
	)

	// Use gorilla/mux for routing
	ro := mux.NewRouter()

	// Set StrictSlash to allow /things/ to automatically redirect to /things
	ro.StrictSlash(true)

	// Use Render for template. Pass in path to templates folder
	// as well as asset helper functions.
	re := render.New(render.Options{
		Layout:     "layouts/layout",
		Extensions: []string{".html"},
	})

	ne.UseHandler(ro)

	return &App{ne, ro, re, db}
}

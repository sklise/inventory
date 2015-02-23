package main

import (
  "fmt"
  "github.com/codegangsta/negroni"
  "github.com/gorilla/context"
  "github.com/gorilla/mux"
  "github.com/unrolled/render"
  // "html/template"
  "log"
  "net/http"
  "time"
)

type Thing struct {
  Id          int64
  AuthorId    int64
  PublisherId int64
  Year        int64
  CreatedAt   time.Time
  UpdatedAt   time.Time
}

type Author struct {
  Id            int64
  CreatedAt     time.Time
  UpdatedAt     time.Time
  Name          string
}

type Publisher struct {
  Id            int64
  CreatedAt     time.Time
  UpdatedAt     time.Time
  Name          string
}

func main() {

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
    Layout: "layouts/layout",
    Extensions: []string{".html"},
  })

  ne.UseHandler(ro)

  ro.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    re.HTML(w, 200, "index", nil)
  }).Methods("Get")

  ro.HandleFunc("/things", func(w http.ResponseWriter, r *http.Request) {
    re.HTML(w, 200, "things/index", nil)
  }).Methods("Get")

  ro.HandleFunc("/things/new", func(w http.ResponseWriter, r *http.Request) {
    re.HTML(w, 200, "things/new", nil)
  }).Methods("Get")

  ro.HandleFunc("/things", func(w http.ResponseWriter, r *http.Request) {
    vars := context.Get(r, "params")
    fmt.Fprintf(w, "%s", vars)
  }).Methods("Post")

  ro.HandleFunc("/things/{id}", func(w http.ResponseWriter, r *http.Request) {
    vars := context.Get(r, "params")

    fmt.Fprintf(w, "%s", vars)
  }).Methods("Get")

  ro.HandleFunc("/things/{id}", func(w http.ResponseWriter, r *http.Request) {
    vars := context.Get(r, "params")
    fmt.Fprintf(w, "%s", vars)
  }).Methods("Put")

  ro.HandleFunc("/things/{id}", func(w http.ResponseWriter, r *http.Request) {
    vars := context.Get(r, "params")
    fmt.Fprintf(w, "ID: %v", vars)
  }).Methods("Delete")

  http.Handle("/", ro)
  http.ListenAndServe(":8080", nil)
}

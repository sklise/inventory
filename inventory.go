package main

import (
  "fmt"
  "github.com/codegangsta/negroni"
  "github.com/gorilla/context"
  "github.com/gorilla/mux"
  _ "github.com/lib/pq"
  "github.com/jinzhu/gorm"
  "github.com/unrolled/render"
  // "html/template"
  // "log"
  "net/http"
  "strings"
  "time"
)

type Thing struct {
  Id          int64
  Year        int64
  Title       string
}

type Author struct {
  Id            int64
  Name          string
}

type Publisher struct {
  Id            int64
  CreatedAt     time.Time
  UpdatedAt     time.Time
  Name          string
}

func main() {
  db, err := gorm.Open("postgres", "dbname=inventory sslmode=disable")
  if err != nil {
    fmt.Printf("Error connecting to database: %v",err)
  }
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
    things := []Thing{}
    db.Find(&things)
    fmt.Printf("%v", len(things))
    for i := 0; i < len(things); i++ {
      fmt.Printf("%v", things[i])
    }
    re.HTML(w, 200, "things/index", things)
  }).Methods("Get")

  ro.HandleFunc("/things/new", func(w http.ResponseWriter, r *http.Request) {
    re.HTML(w, 200, "things/new", nil)
  }).Methods("Get")

  ro.HandleFunc("/things", func(w http.ResponseWriter, r *http.Request) {
    // Parse form values
    err1 := r.ParseForm()
    if err1 != nil {
      fmt.Println("Cannot parse form")
    }
    data := r.PostForm

    thing := Thing{
      Title: strings.Join(data["title"], ""),
    }

    db.Create(&thing)
    fmt.Fprintf(w, "%v", data["title"])
  }).Methods("Post")

  ro.HandleFunc("/things/{id}", func(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    thing := Thing{}
    db.Where("id = ?", vars["id"]).First(&thing)

    if thing.Id == 0 {
      re.HTML(w, 404, "error", "404 Could not find requested thing")
    } else {
      re.HTML(w, 200, "things/show", thing)
    }
  }).Methods("Get")

  ro.HandleFunc("/things/{id}", func(w http.ResponseWriter, r *http.Request) {
    vars := context.Get(r, "params")
    fmt.Fprintf(w, "%s", vars)
  }).Methods("Put")

  ro.HandleFunc("/things/{id}", func(w http.ResponseWriter, r *http.Request) {
    vars := context.Get(r, "params")
    fmt.Fprintf(w, "ID: %v", vars)
  }).Methods("Delete")

  // Index for Authors
  ro.HandleFunc("/authors", func(w http.ResponseWriter, r *http.Request) {
    authors := []Author{}
    db.Find(&authors)

    re.HTML(w, 200, "authors/index", authors)
  }).Methods("Get")

  // New Author Form
  ro.HandleFunc("/authors/new", func(w http.ResponseWriter, r *http.Request) {
    re.HTML(w, 200, "authors/new","")
  }).Methods("Get")

  // Create Author
  ro.HandleFunc("/authors", func(w http.ResponseWriter, r *http.Request) {
    err1 := r.ParseForm()
    if err1 != nil {
      re.HTML(w, 500, "error", 500)
      return
    }

    author := Author {
      Name: r.FormValue("name"),
    }

    db.Create(&author)

    http.Redirect(w,r,"/authors",http.StatusFound)

  }).Methods("Post")

  // Show Author
  ro.HandleFunc("/authors/{id}", func(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    author := Author{}
    if db.Where("id = ?", vars["id"]).First(&author).RecordNotFound() {
      re.HTML(w, 404, "error", "404 Could not find requested author")
      return
    }
    re.HTML(w, 200, "authors/show", author)
  }).Methods("Get")

  http.Handle("/", ro)
  http.ListenAndServe(":8080", nil)
}

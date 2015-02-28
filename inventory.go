package main

import (
  "fmt"
  "github.com/codegangsta/negroni"
  "github.com/gorilla/context"
  "github.com/gorilla/mux"
  _ "github.com/lib/pq"
  "github.com/jinzhu/gorm"
  "github.com/unrolled/render"
  "net/http"
  "os"
  "strconv"
  "time"
)

type Format struct {
  Id          int64
  Name        string
  CreatedAt   time.Time
  UpdatedAt   time.Time
}

type Thing struct {
  Id          int64
  Year        int64
  Title       string
  AuthorId    int64
  CreatedAt   time.Time
  UpdatedAt   time.Time
}

type Author struct {
  Id            int64
  Name          string
  CreatedAt   time.Time
  UpdatedAt   time.Time
}

type Publisher struct {
  Id            int64
  CreatedAt     time.Time
  UpdatedAt     time.Time
  Name          string
}

type AuthorAndThings struct {
  Author Author
  Things []Thing
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
    authors := []Author{}
    db.Find(&authors)
    re.HTML(w, 200, "things/new", authors)
  }).Methods("Get")

  ro.HandleFunc("/things", func(w http.ResponseWriter, r *http.Request) {
    // Parse form values
    err1 := r.ParseForm()
    if err1 != nil {
      fmt.Println("Cannot parse form")
    }

    title := r.FormValue("title")
    year, _ := strconv.ParseInt(r.FormValue("year"), 10, 64)
    author_id, _ := strconv.ParseInt(r.FormValue("author_id"), 10, 64)

    thing := Thing{
      Title: title,
      Year: year,
      AuthorId: author_id,
    }

    err := db.Create(&thing)
    if err != nil {
      re.HTML(w, 500, "error", err)
      return
    }

    things := []Thing{}
    db.Find(&things)
    re.HTML(w, 200, "things/index", things)
  }).Methods("Post")

  // Show Thing
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

  // Modify Thing
  ro.HandleFunc("/things/{id}", func(w http.ResponseWriter, r *http.Request) {
    vars := context.Get(r, "params")
    fmt.Fprintf(w, "%s", vars)
  }).Methods("Put")

  // Delete Thing
  ro.HandleFunc("/things/{id}", func(w http.ResponseWriter, r *http.Request) {
    vars := context.Get(r, "params")
    fmt.Fprintf(w, "ID: %v", vars)
  }).Methods("Delete")

  // Index for Authors
  ro.HandleFunc("/authors", func(w http.ResponseWriter, r *http.Request) {
    authors := []Author{}
    db.Order("LOWER(name) asc").Find(&authors)

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
    things := []Thing{}

    if db.Where("id = ?", vars["id"]).First(&author).RecordNotFound() {
      re.HTML(w, 404, "error", "404 Could not find requested author")
      return
    }

    db.Model(&author).Order("year desc").Related(&things)

    data := AuthorAndThings {
      Author: author,
      Things: things,
    }

    re.HTML(w, 200, "authors/show", data)
  }).Methods("Get")

  // Launch Negroni with all middleware and run on this port.
  ne.Run(":" + os.Getenv("PORT"))
}

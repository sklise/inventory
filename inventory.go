package main

import (
  "fmt"
  "github.com/gorilla/mux"
  "log"
  "net/http"
  "time"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("GET /home")
  fmt.Fprintf(w, "home")
}

func ThingsIndexHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "All things")
}

func ThingsShowHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)

  // Due to how httprouter works, we need to forward to "new" from within the show route
  if vars["id"] == "new" {
    ThingsNewHandler(w,r)
  } else {
    fmt.Fprintf(w, "%s", vars)
  }
}

func ThingsNewHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "new")
}

func ThingsCreateHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  fmt.Fprintf(w, "%s", vars)
}

func ThingsDeleteHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  fmt.Fprintf(w, "ID: %v", vars)
}

func ThingsUpdateHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  fmt.Fprintf(w, "%s", vars)
}

func loggingHandler(next http.Handler) http.Handler {
  fn := func(w http.ResponseWriter, r *http.Request) {
    t1 := time.Now()
    next.ServeHTTP(w, r)
    t2 := time.Now()
    log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
  }

  return http.HandlerFunc(fn)
}

func main() {
  r := mux.NewRouter()

  r.HandleFunc("/", HomeHandler).Methods("GET")
  r.HandleFunc("/things", ThingsIndexHandler).Methods("GET")
  r.HandleFunc("/things", ThingsCreateHandler).Methods("POST")
  r.HandleFunc("/things/new", ThingsNewHandler).Methods("GET")
  r.HandleFunc("/things/{id}", ThingsShowHandler).Methods("GET")
  r.HandleFunc("/things/{id}", ThingsUpdateHandler).Methods("PUT")
  r.HandleFunc("/things/{id}", ThingsDeleteHandler).Methods("DELETE")
  http.Handle("/", r)
  http.ListenAndServe(":8080", nil)
}

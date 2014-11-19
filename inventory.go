package main

import (
  "github.com/gorilla/mux"
  "net/http"
  "fmt"
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
  fmt.Fprintf(w, "%s", vars)
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
  fmt.Fprintf(w, "ID: %s", vars["id"])
}

func ThingsUpdateHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  fmt.Fprintf(w, "%s", vars)
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
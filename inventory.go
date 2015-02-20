package main

import (
  "fmt"
  "github.com/julienschmidt/httprouter"
  "log"
  "net/http"
  "time"
)

func HomeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  fmt.Println("GET /home")
  fmt.Fprintf(w, "home")
}

func ThingsIndexHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  fmt.Fprintf(w, "All things")
}

func ThingsShowHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  vars := ps

  // Due to how httprouter works, we need to forward to "new" from within the show route
  if ps.ByName("id") == "new" {
    ThingsNewHandler(w,r,ps)
  } else {
    fmt.Fprintf(w, "%s", vars)
  }
}

func ThingsNewHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  fmt.Fprintf(w, "new")
}

func ThingsCreateHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  vars := ps
  fmt.Fprintf(w, "%s", vars)
}

func ThingsDeleteHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  vars := ps
  fmt.Fprintf(w, "ID: %v", vars)
}

func ThingsUpdateHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  vars := ps
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
  r := httprouter.New()

  r.GET("/", HomeHandler)

  r.GET("/things", ThingsIndexHandler)
  r.POST("/things", ThingsCreateHandler)
  r.GET("/things/:id", ThingsShowHandler)
  r.PUT("/things/:id", ThingsUpdateHandler)
  r.DELETE("/things/:id", ThingsDeleteHandler)

  http.Handle("/", r)
  http.ListenAndServe(":8080", nil)
}

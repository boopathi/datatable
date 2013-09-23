package main

import (
  "net/http"
  "github.com/gorilla/mux"
  "flag"
  "strconv"
)

var Port int

func main() {
  flag.IntVar(&Port, "p", 4200, "Port Number")
  flag.Parse()

  r := mux.NewRouter()
  r.HandleFunc("/", ListViews)
  r.HandleFunc("/view/{view}", ViewHandler)

  a := r.PathPrefix("/api").Subrouter()
  a.HandleFunc("/put", PutHandler).Methods("POST","PUT")
  a.HandleFunc("/get", GetHandler)

  r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
  http.ListenAndServe(":" + strconv.Itoa(Port), r)

}

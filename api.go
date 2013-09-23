package main

import (
  "net/http"
  "fmt"
  "time"
)

//
// PUT handler - Hadrons
//

func CreateHandler(w http.ResponseWriter, r *http.Request) {
  err := r.ParseForm()
  if err != nil {
    fmt.Println("Parsing error")
    return
  }
  var class, cols string = getval(r.Form["class"]), getval(r.Form["cols"])
  if class == "" || cols == "" {
    fmt.Println("Ignoring creation")
    w.Write([]byte("Invalid class or cols value \n"))
    return
  }
  h := Hadron{
    Class: class,
    Cols: cols,
  }
  err = CreateTable(&h)
  if err != nil {
    fmt.Println(err)
    return
  }
  w.Write([]byte("Table Desc Id = " + h.Id.String() + "\n"))
}

//
// GET handler - Hadron
//

func GetColsHandler(w http.ResponseWriter, r *http.Request) {
  err := r.ParseForm()
  if err != nil {
    fmt.Println("Parsing Error")
    return
  }
  var class string = getval(r.Form["class"])
  if class == "" {
    fmt.Println("Invalid request")
    w.Write([]byte("Invalid class variable\n"))
  }
  h, err := GetTableDesc(class)
  if err != nil {
    fmt.Println(err)
    return
  }
  w.Write([]byte(h.Cols))
}

//
// PUT handler - Quark
//

func PutHandler(w http.ResponseWriter, r *http.Request) {
  err := r.ParseForm()
  if err != nil {
    fmt.Println("Parsing Error")
    return
  }
  var key,class string = getval(r.Form["host"]), getval(r.Form["class"])
  if key == "" || class == "" {
    fmt.Println("Ignoring datapoint")
    return
  }
  q := Quark{
    Ip: r.RemoteAddr,
    Value: getval(r.Form["data"]),
    From: key,
    Class: class,
    Ts: time.Now(),
  }
  err = PutQuark(&q)
  if err != nil {
    fmt.Println(err)
    return
  }

  w.Write([]byte("Insert Id = " + q.Id.String() + "\n"))
}

//
// GET handler - Quark
//

func GetHandler(w http.ResponseWriter, r *http.Request) {
  err := r.ParseForm()
  if err != nil {
    fmt.Println("Error Parsing form", err)
    return
  }
  var key,class string = getval(r.Form["host"]), getval(r.Form["class"])
  if key == "" || class == "" {
    w.Write([]byte("Invalid Request"))
    return
  }
  q, err := GetQuarkByHost(key,class)
  if err != nil {
    fmt.Println("Error fetching Quark", err)
    return
  }
  w.Write([]byte(q.Value))
}


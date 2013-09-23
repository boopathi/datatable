package main

import (
  "net/http"
  "fmt"
  "time"
)

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
  }

  w.Write([]byte("Insert Id = " + q.Id.String() + "\n"))
}
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


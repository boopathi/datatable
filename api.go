package main

import (
  "net/http"
  "fmt"
  "code.google.com/p/couch-go"
  "time"
)
func getval(a []string) string {
  if len(a) > 0 {
    return a[0]
  }
  return ""
}
func PutHandler(w http.ResponseWriter, r *http.Request) {
  err := r.ParseForm()
  if err != nil {
    fmt.Println("Parsing Error")
    return
  }
  var key string = getval(r.Form["host"])
  if key == "" {
    fmt.Println("Ignoring datapoint")
    return
  }
  d := Quark{
    Ip: r.RemoteAddr,
    Value: getval(r.Form["data"]),
    From: getval(r.Form["host"]),
    Class: getval(r.Form["class"]),
    Ts: time.Now(),
  }
  db, err := couch.NewDatabase("localhost","5984", "datatable")
  if err != nil {
    fmt.Println("Error connecting to DB", err)
    return
  }
  tmp := Quark{}
  var id string
  rev, err := db.Retrieve(key,&tmp)
  if err != nil {
    //retrieve failed - so safely insert
    id, rev, err = db.InsertWith(d, key)
    if err != nil {
      fmt.Println("Error inserting", err)
      return
    }
  } else {
    //Retrive and Edit
    id = key
    _, err := db.EditWith(&d, id, rev)
    if err != nil {
      fmt.Println("Error inserting", err)
    }
  }
  w.Write([]byte("Id = " + id + "\nRev = " + rev + "\n"))
}
func GetHandler(w http.ResponseWriter, r *http.Request) {
  err := r.ParseForm()
  if err != nil {
    fmt.Println("Error Parsing form", err)
    return
  }
  var id string = getval(r.Form["host"])
  q, err := GetQuark(id)
  if err != nil {
    fmt.Println("Error fetching Quark", err)
    return
  }
  w.Write([]byte(q.Value))
}

func GetQuark(id string) (Quark, error) {
  db,err := couch.NewDatabase("localhost", "5984", "datatable")
  if err != nil {
    fmt.Println("Error Connecting to DB")
    return Quark{}, err
  }
  data := Quark{}
  _, err = db.Retrieve(id, &data)
  if err != nil {
    fmt.Println("Error retrieving", err)
    return Quark{}, err
  }
  return data, nil
}

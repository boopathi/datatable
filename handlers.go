package main

import (
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
  "strings"
)

func ListViews(w http.ResponseWriter, r *http.Request) {
  C, err := GetCollections()
  if err != nil {
    fmt.Println("Cannot Fetch Collections: ", err)
    return
  }
  page, err := getPage("list", map[string]interface{} {
    "Views": C,
  })
  w.Write(page)
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  class := vars["view"]
  if class == "" {
    w.Write([]byte("Invalid Class"))
    return
  }
  h, err := GetTableDesc(class)
  if err != nil {
    fmt.Println(err)
    return
  }
  Headers := strings.Split(h.Cols, ",")
  qs, err := GetQuarksByClass(vars["view"])
  if err != nil {
    fmt.Println(err)
    return
  }
  var Body [][]string
  for _, v := range qs {
    lines := strings.Split(v.Value, "\n")
    for _, line := range lines {
      if line == "" { continue }
      Body = append(Body,strings.Split(line, ","))
    }
  }
  page, err := getPage("datatable", map[string]interface{} {
    "Headers": Headers,
    "Body": Body,
  })
  if err != nil {
    fmt.Println(err)
    return
  }
  w.Write(page)
}

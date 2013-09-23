package main

import (
  "fmt"
  "net/http"
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
  a := ViewTableData{}
  page, err := getPage("datatable",a)
  if err != nil {
    fmt.Println(err)
    return
  }
  w.Write(page)
}

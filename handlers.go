package main

import (
  "fmt"
  "net/http"
)

func ViewDataTable(w http.ResponseWriter, r *http.Request) {
  a := ViewTableData{}
  page, err := getPage("datatable",a)
  if err != nil {
    fmt.Println(err)
    return
  }
  w.Write(page)
}

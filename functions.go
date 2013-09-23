package main

import (
  "bytes"
  "html/template"
  "os"
  "code.google.com/p/couch-go"
  "fmt"
)

func getval(a []string, d ...string) string {
  if len(a) > 0 {
    return a[0]
  }
  if len(d) > 0 {
    return d[0]
  }
  return ""
}

func GetQuarkById(id string) (Quark, error) {
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

func GetQuarksByClass(class string) ([]Quark, error) {
  db, err := couch.NewDatabase("localhost", "5984", "datatable")
  if err != nil {
    fmt.Println("Error Connecting to DB")
    return nil, err
  }
}

func parseTemplate(file string, data interface{}) ([]byte, error) {
  var buf bytes.Buffer
  t, err := template.ParseFiles(file)
  if err != nil {
    return nil, err
  }
  err = t.Execute(&buf, data)
  if err != nil {
    return nil, err
  }
  return buf.Bytes(), nil
}

func getPage(tmpl string, data interface{}) ([]byte, error) {
  filename := "templates/" + tmpl + ".html"
  if _,err := os.Stat(filename); err != nil {
    return nil, err
  }
  return parseTemplate(filename, data)
}


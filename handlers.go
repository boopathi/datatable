package main

import (
  "fmt"
  "bytes"
  "html/template"
  "net/http"
  "github.com/gorilla/mux"
  "labix.org/v2/mgo/bson"
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

func DT_GetHeaders(class string, h chan<- []string, e chan<- error) {
  ha, err := GetTableDesc(class)
  if err != nil {
    e<- err
    return
  }
  Headers := strings.Split(ha.Cols, ",")
  h<- Headers
}

func ParseQuark(q Quark, b chan<- [][]string) {
  var Body [][]string
  lines := strings.Split(q.Value, "\n")
  for _, line := range lines {
    if line == "" { continue }
    Body = append(Body, strings.Split(line,","))
  }
  b<- Body
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
  var err error
  vars := mux.Vars(r)
  class := vars["view"]
  if class == "" {
    w.Write([]byte("Invalid Class"))
    return
  }

  //Throw Page initial content first
  page,err := getPage("header", nil)
  if err != nil { fmt.Println(err); return }
  w.Write(page)

  head := make(chan []string)
  ehead := make(chan error)
  go DT_GetHeaders(class,head,ehead)

  c := DB.Db.C(class)
  var q Quark
  body := make(chan [][]string)
  iter := c.Find(bson.M{}).Iter()
  count := 0
  for iter.Next(&q) {
    count = count + 1
    go ParseQuark(q, body)
  }

  var buf bytes.Buffer
  //wait to throw head
  select {
  case Headers := <-head :
    t, err := template.ParseFiles(Config.TmplDir + "/dt_header.html")
    if err != nil { fmt.Println(err); return}
    t.Execute(&buf, Headers)
    w.Write(buf.Bytes())

  case err = <-ehead:
    fmt.Println(err)
    return
  }

  //Now the middle portion
  page, err = getPage("middle", nil)
  if err != nil { fmt.Println(err); return }
  w.Write(page)

  //Now Throw Body
  for i:=0;i<count;i++ {
    Body := <-body
    t, err := template.ParseFiles(Config.TmplDir + "/dt_body.html")
    if err != nil { fmt.Println(err); return }
    t.Execute(&buf, Body)
    w.Write(buf.Bytes())
  }

  page,err = getPage("footer", nil)
  if err != nil { fmt.Println(err); return }
  w.Write(page)

}

func ViewHandler_Old(w http.ResponseWriter, r *http.Request) {
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

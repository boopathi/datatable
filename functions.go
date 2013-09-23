package main

import (
  "bytes"
  "html/template"
  "os"
  "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson"
  "fmt"
  "errors"
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

func ConnectDB() (*mgo.Database, chan<- bool, error) {
  session, err := mgo.Dial("localhost")
  if err != nil {
    fmt.Println("Cannot Dial to mongodb")
    return nil, nil, err
  }
  // TODO: Check if we can implement "TIME OUT, MOVE ON" pattern
  c := make(chan bool)
  go func() {
    <-c
    session.Close()
  }()
  d := session.DB("datatable")
  return d, c, nil
}

//
// Getters - Collections
//

func GetCollections() ([]string, error) {
  session, err := mgo.Dial("localhost")
  if err != nil { return nil,err }
  defer session.Close()
  db := session.DB("datatable")
  c, err := db.CollectionNames()
  var r []string
  for _, v := range c {
    if len(v)>3 && v[0:2] == "__" && v[len(v)-2:] == "__" { continue }
    if v == "system.indexes" { continue }
    r = append(r,v)
  }
  return r,err
}

//
// Getters and Setters - Hadrons
//

func CreateTable(h *Hadron) error{
  session, err := mgo.Dial("localhost")
  if err != nil { return err }
  defer session.Close()
  h.Id = bson.NewObjectId()
  c := session.DB("datatable").C("__colstable__")
  _, err = c.Upsert(bson.M{"class": h.Class}, &h)
  if err != nil { return err }
  return nil
}

func GetTableDesc(class string) (Hadron, error) {
  session, err := mgo.Dial("localhost")
  if err != nil { return Hadron{}, err }
  defer session.Close()
  c := session.DB("datatable").C("__colstable__")
  var results []Hadron
  err = c.Find(bson.M{"class": class}).All(&results)
  if err != nil { return Hadron{}, err }
  if len(results) < 1 {
    return Hadron{}, errors.New("GetTableDesc: Fetched None")
  }
  return results[0], nil
}

//
// Getters and Setters - Quarks
//

func PutQuark(q *Quark) error {
  session, err := mgo.Dial("localhost")
  if err != nil { return err }
  defer session.Close()
  q.Id = bson.NewObjectId()
  c := session.DB("datatable").C(q.Class)
  _, err = c.Upsert(bson.M{"from":q.From, "class":q.Class}, &q)
  if err != nil { return err }
  return nil
}

func GetQuarkByHost(host, class string) (Quark, error) {
  session, err := mgo.Dial("localhost")
  if err != nil { return Quark{}, err }
  defer session.Close()
  c :=  session.DB("datatable").C(class)

  var results []Quark
  err = c.Find(bson.M{"from": host, "class":class}).All(&results)
  if err != nil { return Quark{}, err }

  if len(results) < 1 {
    return Quark{}, errors.New("GetQuarksById: Fetched None")
  }
  return results[0], nil
}

func GetQuarksByClass(class string) ([]Quark, error) {
  session, err := mgo.Dial("localhost")
  if err != nil { return nil, err }
  defer session.Close()
  c := session.DB("datatable").C(class)

  var results []Quark
  err = c.Find(bson.M{}).All(&results)
  if err != nil { return nil, err }

  return results, nil
}

//
// Template Functions
//

func parseTemplate(file string, data interface{}) ([]byte, error) {
  var buf bytes.Buffer
  t, err := template.ParseFiles(file)
  if err != nil { return nil, err }
  err = t.Execute(&buf, data)
  if err != nil { return nil, err }
  return buf.Bytes(), nil
}

func getPage(tmpl string, data interface{}) ([]byte, error) {
  filename := "templates/" + tmpl + ".html"
  if _,err := os.Stat(filename); err != nil { return nil, err }
  return parseTemplate(filename, data)
}


package main

import (
	"bytes"
	"errors"
	"html/template"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
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

func ConnectDB() error {
	var err error
	DB.Session, err = mgo.Dial(Config.DBHost)
	if err != nil {
		return err
	}
	DB.Db = DB.Session.DB(Config.DBName)
	return nil
}

//
// Getters - Collections
//

func GetCollections() ([]string, error) {
	c, err := DB.Db.CollectionNames()
	var r []string
	for _, v := range c {
		if len(v) > 3 && v[0:2] == "__" && v[len(v)-2:] == "__" {
			continue
		}
		if v == "system.indexes" {
			continue
		}
		r = append(r, v)
	}
	return r, err
}

//
// Getters and Setters - Hadrons
//

func CreateTable(h *Hadron) error {
	h.Id = bson.NewObjectId()
	c := DB.Db.C("__colstable__")
	_, err := c.Upsert(bson.M{"class": h.Class}, &h)
	if err != nil {
		return err
	}
	return nil
}

func GetTableDesc(class string) (Hadron, error) {
	c := DB.Db.C("__colstable__")
	var results []Hadron
	err := c.Find(bson.M{"class": class}).All(&results)
	if err != nil {
		return Hadron{}, err
	}
	if len(results) < 1 {
		return Hadron{}, errors.New("GetTableDesc: Fetched None")
	}
	return results[0], nil
}

//
// Getters and Setters - Quarks
//

func PutQuark(q *Quark) error {
	q.Id = bson.NewObjectId()
	c := DB.Db.C(q.Class)
	_, err := c.Upsert(bson.M{"from": q.From, "class": q.Class}, &q)
	if err != nil {
		return err
	}
	return nil
}

func GetQuarkByHost(host, class string) (Quark, error) {
	c := DB.Db.C(class)

	var results []Quark
	err := c.Find(bson.M{"from": host, "class": class}).All(&results)
	if err != nil {
		return Quark{}, err
	}

	if len(results) < 1 {
		return Quark{}, errors.New("GetQuarksById: Fetched None")
	}
	return results[0], nil
}

func GetQuarksByClass(class string) ([]Quark, error) {
	c := DB.Db.C(class)
	var results []Quark
	err := c.Find(bson.M{}).All(&results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

//
// Template Functions
//

func getTmplName(tmpl string) string {
	return Config.TmplDir + "/" + tmpl + ".html"
}

func getTemplate(tmpl string) (*template.Template, error) {
	t, err := template.ParseFiles(getTmplName(tmpl))
	return t, err
}

func parseTemplate(tmpl *template.Template, data interface{}) []byte {
	var buf bytes.Buffer
	tmpl.Execute(&buf, data)
	return buf.Bytes()
}

func getPage(tmpl string, data interface{}) ([]byte, error) {
	t, err := getTemplate(tmpl)
	if err != nil {
		return nil, err
	}
	return parseTemplate(t, data), nil
}

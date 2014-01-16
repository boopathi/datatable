package main

import (
	"github.com/gorilla/mux"
	"labix.org/v2/mgo/bson"
	"log"
	"net/http"
	"strings"
)

func ListViews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("X-Powered-By", "datatable.go")

	C, err := GetCollections()
	if err != nil {
		log.Println("Cannot Fetch Collections: ", err)
		return
	}
	page, err := getPage("list", map[string]interface{}{
		"Views": C,
	})
	w.Write(page)
	Log(r)
}

func ParseQuark(q Quark, b chan<- [][]string) {
	var Body [][]string
	lines := strings.Split(q.Value, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		Body = append(Body, strings.Split(line, ","))
	}
	b <- Body
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	class := vars["view"]
	if class == "" {
		w.Write([]byte("Invalid Class"))
		return
	}

	//Set HTTP headers
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("X-Powered-By", "datatable.go")

	//Get Header
	h, err := GetTableDesc(class)
	if err != nil {
		log.Println(err)
		return
	}
	Headers := strings.Split(h.Cols, ",")

	//Throw Page initial content first
	page, err := getPage("header", Headers)
	if err != nil {
		log.Println(err)
		return
	}
	w.Write(page)

	//Make the body
	c := DB.Db.C(class)
	var q Quark
	body := make(chan [][]string)
	iter := c.Find(bson.M{}).Iter()
	count := 0
	for iter.Next(&q) {
		count = count + 1
		go ParseQuark(q, body)
	}

	//Now Throw Body
	tmpl, err := getTemplate("dt_body")
	if err != nil {
		log.Println(err)
		return
	}
	for i := 0; i < count; i++ {
		Body := <-body
		w.Write(parseTemplate(tmpl, Body))
	}

	//And finally the footer
	page, err = getPage("footer", nil)
	if err != nil {
		log.Println(err)
		return
	}
	w.Write(page)
	Log(r)
}

//
// Deprecated
//

func ViewHandler_Old(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	class := vars["view"]
	if class == "" {
		w.Write([]byte("Invalid Class"))
		return
	}
	h, err := GetTableDesc(class)
	if err != nil {
		log.Println(err)
		return
	}
	Headers := strings.Split(h.Cols, ",")
	qs, err := GetQuarksByClass(vars["view"])
	if err != nil {
		log.Println(err)
		return
	}
	var Body [][]string
	for _, v := range qs {
		lines := strings.Split(v.Value, "\n")
		for _, line := range lines {
			if line == "" {
				continue
			}
			Body = append(Body, strings.Split(line, ","))
		}
	}
	page, err := getPage("datatable", map[string]interface{}{
		"Headers": Headers,
		"Body":    Body,
	})
	if err != nil {
		log.Println(err)
		return
	}
	w.Write(page)
}

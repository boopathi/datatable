package main

import (
	"bitbucket.org/kardianos/osext"
	"flag"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"path"
	"runtime"
	"strconv"
	"time"
)

var Config Conf
var DB Database

func main() {
	//Set runtime.GOMAXPROCS
	runtime.GOMAXPROCS(runtime.NumCPU())

	current, _ := osext.Executable()
	current = path.Dir(current)

	flag.IntVar(&Config.Port, "port", 4200, "Server port Number")
	flag.StringVar(&Config.DBHost, "dbhost", "localhost", "MongoDB Host")
	flag.IntVar(&Config.DBPort, "dbport", 27017, "MongoDB Port")
	flag.StringVar(&Config.DBUser, "dbuser", "", "MongoDB User")
	flag.StringVar(&Config.DBPass, "dbpass", "", "MongoDB Password")
	flag.StringVar(&Config.DBName, "dbname", "datatable", "MongoDB Database Name")
	flag.StringVar(&Config.StaticDir, "staticdir", current+"/static", "Static Directory Path - Absolute")
	flag.StringVar(&Config.TmplDir, "tmpldir", current+"/templates", "Template Directory Path - Absolute")
	flag.Parse()

	//Connect to DB
	err := ConnectDB()
	if err != nil {
		log.Println("Error Connecting to DB", err)
		return
	}
	defer DB.Session.Close()

	//Refresh DB Connection every 5 minutes
	//This is because Mongo timesout connection after 10 mins
	ticker := time.NewTicker(5 * time.Minute)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				err = DB.Session.Ping()
				if err == nil {
					log.Println("Refreshing MongoDB Session.")
					DB.Session.Refresh()
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	defer close(quit)

	r := mux.NewRouter()
	r.HandleFunc("/", ListViews)
	r.HandleFunc("/view/{view}", ViewHandler)

	a := r.PathPrefix("/api").Subrouter()
	a.HandleFunc("/create", CreateHandler).Methods("POST", "PUT")
	a.HandleFunc("/cols", GetColsHandler)
	a.HandleFunc("/put", PutHandler).Methods("POST", "PUT")
	a.HandleFunc("/get", GetHandler)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir(Config.StaticDir + "/")))
	err = http.ListenAndServe(":"+strconv.Itoa(Config.Port), r)
	if err != nil {
		log.Println(err)
		return
	}
}

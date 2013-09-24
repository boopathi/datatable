package main

import (
  "net/http"
  "github.com/gorilla/mux"
  "flag"
  "fmt"
  "strconv"
  "io/ioutil"
  "encoding/json"
  "time"
)

var Config Conf
var DB Database

func main() {
  var conffile string
  flag.StringVar(&conffile, "config", "./datatable.json", "Config File")
  flag.Parse()

  c, err := ioutil.ReadFile(conffile)
  if err != nil {
    fmt.Println("Error Reading config \n", err, "\n")
    return
  }
  json.Unmarshal(c, &Config)

  // Validate
  if Config.Port == 0 { Config.Port = 4200 }
  if Config.DBHost == "" { Config.DBHost = "localhost" }
  if Config.DBName == "" { Config.DBName = "datatable" }
  if Config.StaticDir == "" { Config.StaticDir = "./static" }
  if Config.TmplDir == "" { Config.TmplDir = "./templates" }

  //Connect to DB
  err = ConnectDB()
  if err != nil {
    fmt.Println("Error Connecting to DB")
    return
  }
  defer DB.Session.Close()

  //Refresh DB Connection every 5 minutes
  //This is because Mongo timesout connection after 10 mins
  ticker := time.NewTicker( 5 * time.Minute )
  quit := make(chan struct{})
  go func() {
    for {
      select {
      case <-ticker.C:
        err = DB.Session.Ping()
        if err == nil {
          fmt.Println(time.Now(), "Refreshing MongoDB Session.")
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
  a.HandleFunc("/put", PutHandler).Methods("POST","PUT")
  a.HandleFunc("/get", GetHandler)

  r.PathPrefix("/").Handler(http.FileServer(http.Dir(Config.StaticDir + "/")))
  http.ListenAndServe(":" + strconv.Itoa(Config.Port), r)

}

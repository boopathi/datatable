package main

import "time"

// Data from 1 Node
type Quark struct {
  Ip          string
  Value       string
  From        string
  Class       string
  Type        string
  Ts          time.Time
  ClassView   string         `json: "_view/class"`
}

type Conf struct {
  Port        int
  DBHost      string
  DBPort      int
  DBData      string
}

type ViewTableData struct {
  Title       string
  Body        string
}

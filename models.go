package main

import "time"

// Data from 1 Node
type Quark struct {
  Ip      string
  Value   string
  From    string
  Class   string
  Type    string
  Ts      time.Time
}

type ViewTableData struct {
  Title   string
  Body    string
}

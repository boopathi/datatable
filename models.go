package main

import "time"

// Data from 1 Node
type Quark struct {
  Ip      string
  Value   string
  From    string
  Class   string
  Ts      time.Time
}

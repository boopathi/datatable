package main

import "time"
import "labix.org/v2/mgo/bson"

// Data from 1 Node
type Quark struct {
  Id          bson.ObjectId     `bson: "_id,omitempty"`
  Ip          string            `bson: "ip"`
  Value       string            `bson: "value"`
  From        string            `bson: "from"`
  Class       string            `bson: "class"`
  Ts          time.Time         `bson: "ts"`
}

// Table Description
type Hadron struct {
  Id          bson.ObjectId     `bson: "_id,omitempty"`
  Class       string            `bson: "class"`
  Cols        string            `bson: "cols"`
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

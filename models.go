package main

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

type Database struct {
	Session *mgo.Session
	Db      *mgo.Database
}

// Data from 1 Node
type Quark struct {
	Id    bson.ObjectId `bson: "_id,omitempty"`
	Ip    string        `bson: "ip"`
	Value string        `bson: "value"`
	From  string        `bson: "from"`
	Class string        `bson: "class"`
	Ts    time.Time     `bson: "ts"`
}

// Table Description
type Hadron struct {
	Id    bson.ObjectId `bson: "_id,omitempty"`
	Class string        `bson: "class"`
	Cols  string        `bson: "cols"`
}

type Conf struct {
	Port      int    `json: "port"`
	DBHost    string `json: "db_host"`
	DBPort    int    `json: "db_port"`
	DBName    string `json: "db_name"`
	DBUser    string `json: "db_user"`
	DBPass    string `json: "db_pass"`
	StaticDir string `json: "static_dir"`
	TmplDir   string `json: "templates_dir"`
}

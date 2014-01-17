datatable
=========

Collect Tabular data from various endpoints and display them using datatable.js

## Introduction

The purpose of this tool is to collect the `current State` of applications across many servers in *tabular data format*. 

### Terminology

Often used when talking to the `/api` endpoint

+ *Class* - A view or a Table Heading. eg: apache\_domains, dns\_zones
+ *Host* - Short host name.
+ *Cols* - A list of Comma(,) separated values defining the column names for a particular class.
+ *data* - CSV data (Rows are separated by new-line character(\n), Row values by comma(,)).

## LICENSE

MIT

## Installation

+ `go get github.com/boopathi/datatable`
+ `go install`
+ Install `mongodb` and configure `datatable.json`
+ `$GOPATH/bin/datatable -config /path/to/datatable.json`

## Installation - Docker container

+ `docker pull boopathi/mongodb`
+ `docker pull boopathi/datatable`
+ `docker run -p 27017:27017 -v /var/lib/mongodb:/data/db -name mongodb boopathi/mongodb`
+ `docker run -p 4200:4200 -name datatable -link mongodb:db boopathi/datatable`

## Configuration

Usage of `./datatable`:
  -dbhost="localhost": MongoDB Host
  -dbname="datatable": MongoDB Database Name
  -dbpass="": MongoDB Password
  -dbport=27017: MongoDB Port
  -dbuser="": MongoDB User
  -port=4200: Server port Number
  -staticdir="/go/src/github.com/boopathi/datatable/static": Static Directory Path - Absolute
  -tmpldir="/go/src/github.com/boopathi/datatable/templates": Template Directory Path - Absolute

## Distribution

The following options are available for packaging the compiled application. `deb` and `rpm` requires `fpm`.

+ `make tar`
+ `make rpm`
+ `make deb`

## Sending and Receiving Data

### Creating a Class

`PUT /api/create`

+ class = classname
+ cols = "col1,col2,col3,col4"

### Adding Data

`PUT /api/put`

+ class = classname
+ host = hostname
+ data = data

### Get Columns for a particular Class

`GET /api/cols`

+ class = classname

### Get Data for a particular Host under a Class

`GET /api/get`

+ class = classname
+ host = hostname

## Dockers

[https://index.docker.io/u/boopathi/datatable/](https://index.docker.io/u/boopathi/datatable/)

## Contributors

+ Boopathi Rajaa <boopathi>

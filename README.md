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

## Configuration

Sample Configuration file: `datatable.json`. Go through the file. Pretty much self explanatory

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

## Contributors

+ Boopathi Rajaa <boopathi>

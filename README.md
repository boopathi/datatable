datatable
=========

Collect Tabular data from various endpoints and display them using datatable.js

# LICENSE

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

PUT /api/create

+ class = classname
+ cols = "col1,col2,col3,col4"

### Adding Data

PUT /api/put

+ class = classname
+ host = hostname
+ data = data

### Get Columns for a particular Class

GET /api/cols

+ class = classname

### Get Data for a particular Host under a Class

GET /api/get

+ class = classname
+ host = hostname

## Contributors

+ Boopathi Rajaa <boopathi>

datatable
=========

Collect Tabular data from various endpoints and display them using datatable.js

# LICENSE

MIT

## Installation

+ `go get github.com/boopathi/datatable`
+ `make deb` or `make rpm`
+ `sudo service datatable start #init script`

## Configuration

Sample Configuration file: `datatable.json`. Go through the file. Pretty much self explanatory

## Custom Usage

`$GOPATH/bin/datatable -config /path/to/datatable.json`

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

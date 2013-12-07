docker-version  0.6.5
from  ubuntu:12.04
maintainer  Boopathi Rajaa <me@boopathi.in>

# Mercurial
run echo 'deb http://ppa.launchpad.net/mercurial-ppa/releases/ubuntu precise main' > /etc/apt/sources.list.d/mercurial.list
run echo 'deb-src http://ppa.launchpad.net/mercurial-ppa/releases/ubuntu precise main' >> /etc/apt/sources.list.d/mercurial.list
run apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 323293EE

# Install dependencies
run apt-get update
run apt-get install -y curl git bzr mercurial

# Install Go
run curl -s https://go.googlecode.com/files/go1.2.linux-amd64.tar.gz | tar -v -C /usr/local/ -xz
env PATH  /usr/local/go/bin:/usr/local/bin:/usr/local/sbin:/usr/bin:/usr/sbin:/bin:/sbin
env GOPATH  /go
env GOROOT  /usr/local/go

# Expose port
expose 4200

workdir /go/src/github.com/boopathi/datatable

add . /go/src/github.com/boopathi/datatable

# Install application
run go get
run go build

run echo '\
{ \
  "port": 4200,\
  "db_host": $DB_PORT_27017_TCP_ADDR,\
  "db_port": $DB_PORT_27017_TCP_PORT,\
  "db_name": "datatable",\
  "static_dir": "./static",\
  "templates_dir": "./templates"\
}' > ./datatable.json

entrypoint ./datatablae

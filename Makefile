
pth=src/github.com/boopathi/datatable

all:
	go install

tar:
	cp ${GOPATH}/bin/datatable ./
	tar -cvzf datatable.tar.gz ./datatable ./datatable.json ./static/ ./templates/
	rm -f ./datatable

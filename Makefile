
all:
	go install

tar:
	go build
	tar -cvzf datatable.tar.gz ./datatable ./datatable.json ./static/ ./templates/
	rm -f ./datatable

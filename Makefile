NAME=datatable
VER=0.1.3
PREFIX=/opt/datatable
LICENSE=MIT
AUTHOR="me@boopathi.in"
DESC="Collect Tabular data from various endpoints and display them using datatable.js"
URL="https://github.com/boopathi/datatable"

all:
	go install

tar:
	go build
	tar -cvzf datatable.tar.gz ./datatable ./datatable.json ./static/ ./templates/ ./README.md ./LICENSE
	rm -f ./datatable

deb:
	go build
	fpm -s dir -t deb -n ${NAME} -v ${VER} \
		--prefix ${PREFIX} \
		--license ${LICENSE} \
		--provides datatable \
		-m ${AUTHOR} \
		--description ${DESC} \
		--url ${URL} \
		./datatable ./datatable.json ./static/ ./templates/ ./README.md ./LICENSE
	rm -f ./datatable


rpm:
	go build
	fpm -s dir -t rpm -n ${NAME} -v ${VER} \
		--prefix ${PREFIX} \
		--license ${LICENSE} \
		--provides datatable \
		-m ${AUTHOR} \
		--description ${DESC} \
		--url ${URL} \
		./datatable ./datatable.json ./static/ ./templates/ ./README.md ./LICENSE
	rm -f ./datatable

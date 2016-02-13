GOPATH := ${PWD}/_vendor:${GOPATH}
export GOPATH

.PHONY: dependencies definitions bootstrap

default: bootstrap

build:
	go build -v

dependencies:
	go get ./...

definitions:
	protoc -I definitions --go_out lib/definitions/credence definitions/*.proto

bootstrap: dependencies definitions build

install:
	INSTALL_ROOT="/usr/local/"
	SHARE_DIR="${INSTALL_ROOT}/share/credence/"
	mkdir -p ${SHARE_DIR}
	cp -R {htdocs,templates} ${SHARE_DIR}
	cp ./credence "${INSTALL_ROOT}/bin/"

clean:
	rm ./credence
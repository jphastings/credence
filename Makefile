GOPATH := ${PWD}/_vendor:${GOPATH}
export GOPATH

.PHONY: definitions dependencies bootstrap

default: build

build:
	go build -v

dependencies:
	go get ./...

definitions:
	protoc -I definitions --go_out lib/definitions/credence definitions/*.proto

bootstrap: dependencies definitions build
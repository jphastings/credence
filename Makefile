GOPATH := ${PWD}/_vendor:${GOPATH}
export GOPATH

.PHONY: dependencies definitions bootstrap

default: build

build:
	go build -v

dependencies:
	go get ./...

definitions:
	protoc -I definitions --go_out lib/definitions/credence definitions/*.proto

bootstrap: dependencies definitions build
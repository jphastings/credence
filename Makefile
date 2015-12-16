GOPATH := ${PWD}/_vendor:${GOPATH}
export GOPATH

default: build

build:
	go build -v

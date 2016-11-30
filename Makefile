all: build

build:
	GOPATH=`pwd` go build -o bin/eye eye

all: dep build

dep:
	go get github.com/go-yaml/yaml
	go get github.com/jcelliott/lumber
	go get github.com/garyburd/redigo/redis


build:
	go build -o bin/eye eye

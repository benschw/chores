SHELL=/bin/bash

default: build
clean:
	rm -rf build

deps:
	go get -u -t -v ./...

test:
	go test ./...

build: 
	mkdir -p build
	go build -o build/chores

db:
	mysql -u admin -pchangeme -h 172.20.20.1 -e 'CREATE DATABASE Chores'

.PHONY: build

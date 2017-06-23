SHELL=/bin/bash

default: build

deps:
	go get -u -t -v ./...
	go get github.com/golang/lint/golint
	go get github.com/mitchellh/gox

test:
	go test ./...

build:
	gox -output "chores" -osarch="linux/amd64"

docker:
	docker build -t benschw/chores .

publish:
	docker push benschw/chores

release: test build docker publish

db:
	mysql -u admin -pchangeme -h 172.20.20.1 -e 'CREATE DATABASE Chores'

.PHONY: build

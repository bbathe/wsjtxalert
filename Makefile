package := $(shell basename `pwd`)

.PHONY: default get codetest build setup test fmt lint vet

default: fmt codetest

get:
	go get -v ./...
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s -- -b $(shell go env GOPATH)/bin v1.20.0

codetest: lint vet test

build:
	mkdir -p target
	rm -f target/*
	GOOS=windows GOARCH=amd64 go build -v -o target
	cp $(package).yaml target/

test:
	go test

fmt:
	go fmt ./...

lint:
	$(shell go env GOPATH)/bin/golangci-lint run --fix

vet:
	go vet -all .
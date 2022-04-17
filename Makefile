SHELL := /bin/bash

build:
	go build -o main main.go

build-windows:
  GOOS=windows GOARCH=amd64 go build -o main.go

debug:
	go run main.go

run: build
	./main

test:
	go test -race -shuffle on ./...
	go test -coverprofile cover.out
	go tool cover -html=cover.out

clean:
	rm main
	rm cover.out

setup:
	export GOPATH="/workspaces/Go"

update:
	go mod vendor
	go mod tidy

lint:
	golangci-lint  run ./...
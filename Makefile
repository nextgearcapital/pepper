.PHONY: all fmt deps test build

all: fmt deps test build

deps:
	go get ./...

fmt:
	go fmt `go list ./...`

build:
	gox -osarch="linux/amd64"

test:
	go get -t -v ./...
	go tool vet .
	go test -v -race ./...

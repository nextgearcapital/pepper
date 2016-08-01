.PHONY: all fmt deps test build

all: fmt deps test build

fmt:
	go fmt `go list ./...`

deps:
	go get github.com/mitchellh/gox
	go get -t -v ./...

test:
	go get -t -v ./...
	go tool vet .
	go test -v -race ./...

build:
	gox -osarch="linux/amd64"

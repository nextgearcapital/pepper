.PHONY: all fmt deps build

all: fmt deps build

deps:
	go get ./...

fmt:
	go fmt `go list ./...`

build:
	gox -osarch="linux/amd64"

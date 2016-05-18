.PHONY: all fmt build

all: fmt build

fmt:
	go fmt `go list ./...`

build:
	gox -osarch="linux/amd64"

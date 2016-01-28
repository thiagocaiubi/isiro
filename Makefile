.PHONY: all clean build test

all: clean build test

clean:
	@go clean -i ./...

build:
	@go build ./...

test:
	@go test -v ./...

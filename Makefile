ORGANIZATION := thiagocaiubi
CONTAINER := isiro

.PHONY: all
all: clean get build test

.PHONY: build
build:
	go build ./...

.PHONY: clean
clean:
	go clean -i ./...

.PHONY: docker-build
docker-build:
	docker build -t $(ORGANIZATION)/$(CONTAINER) .

.PHONY: docker
docker: docker-build
	docker run --rm \
		-v $(PWD):/go/src/$(CONTAINER) \
		--name $(CONTAINER) \
		$(ORGANIZATION)/$(CONTAINER)

.PHONY: get
get:
	go get -v -d ./...

.PHONY: test
test:
	go test -v ./...

.PHONY: test-race
test-race:
	go test -race -v ./...

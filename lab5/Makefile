
all:
	docker compose run --rm dev bash

.PHONY: test
test: test-fast build
	@cd tests && pytest

.PHONY: test-fast
test-fast:
	@go test -v ./...

.PHONY: build
build:
	@go build -o lab5 cmd/lab5/lab5.go

# Start targets makefile in docker environment
docker-impl-%:
	docker compose run --rm dev make $*

.PHONY: docker-test
docker-test: docker-impl-test

.PHONY: docker-test-fast
docker-test-fast: docker-impl-test-fast

.PHONY: docker-build
docker-build: docker-impl-build

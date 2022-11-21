
all:
	docker compose run --rm dev bash

.PHONY: test
test: test-fast build
	@cd tests && pytest --verbose

.PHONY: test-fast
test-fast:
	@go test -v ./...

.PHONY: build
build:
	@ CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o lab5 cmd/lab5/lab5.go

.PHONY: client
client:
	@ go run cmd/client/client.go

# Start targets makefile in docker environment
docker-impl-%:
	docker compose run --rm dev make $*

.PHONY: docker-test
docker-test: docker-impl-test

.PHONY: docker-test-fast
docker-test-fast: docker-impl-test-fast

.PHONY: docker-build
docker-build: docker-impl-build

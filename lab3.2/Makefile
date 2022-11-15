
.PHONY: test
test: test-fast
	@cd tests && pytest

.PHONY: test-fast
test-fast:
	@echo "unimplemented"

# Start targets makefile in docker environment
docker-impl-%:
	docker compose run --rm dev -c make $*

.PHONY: docker-test
docker-test: docker-impl-test

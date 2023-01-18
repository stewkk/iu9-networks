
all: sync

sync:
	go run cmd/sync/sync.go

async:
	go run cmd/async/async.go

auth:
	go build -tags pam cmd/auth/auth.go && sudo ./auth

.PHONY: all sync async auth

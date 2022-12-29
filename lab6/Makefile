
all:
	go run cmd/lab6/lab6.go

server:
	go run cmd/server/server.go

rss:
	go run cmd/rss-client/rss-client.go

list:
	cat tmp/rss/* | jq '.[] .Id'

.PHONY: all server rss

server:
	go run cmd/server/main.go -port 8080
test:
	go test -cover -race ./...
build:
	go build ./cmd/server/main.go

.PHONY: gen clean server test client

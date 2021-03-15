server:
	go run cmd/server/main.go -port 8080
test:
	go test -cover -race ./...
gen:
	export PATH="$PATH:$(go env GOPATH)/bin"
	protoc --go_out=./pb --go_opt=paths=source_relative \
    --go-grpc_out=./pb --go-grpc_opt=paths=source_relative \
    --proto_path=proto proto/*.proto -I.
clean:
	rm pb/*.go
build:
	go build ./cmd/server/main.go

.PHONY: gen clean server test client

FROM golang:1.16-alpine AS build

RUN apk add --no-cache git
RUN apk add build-base

WORKDIR /cemiac

COPY . .

RUN go mod download

RUN go test -cover -race ./...

RUN go build -o ./temp/mailServer cmd/server/main.go

FROM alpine:3.9 AS bin

COPY --from=build /cemiac/temp/mailServer /app/mailServer

ENTRYPOINT ["/app/mailServer","-port=8080"]

EXPOSE 8080

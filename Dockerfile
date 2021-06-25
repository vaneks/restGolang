FROM golang:latest

WORKDIR /app

COPY ./ /app


ENTRYPOINT go run main.go
FROM golang:1.15.6-alpine3.12

ENV ENV GO11MODULE=on

WORKDIR /app

COPY . /app

RUN apk update && \
    apk add --no-cache git && \
    go mod download

EXPOSE 8080

FROM golang:1.17 AS builder

WORKDIR /app

COPY ./opaque-client ./opaque-client
COPY ./opaque-service-protos ./opaque-service-protos

WORKDIR /app/opaque-client

RUN go mod download 

RUN go build -o ./opaque-client-binary
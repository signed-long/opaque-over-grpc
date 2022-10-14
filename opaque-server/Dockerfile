FROM golang:1.17 AS builder

WORKDIR /app

COPY ./opaque-server ./opaque-server
COPY ./opaque-service-protos ./opaque-service-protos

WORKDIR /app/opaque-server

RUN go mod download 

RUN go build -o ./opaque-server-binary

CMD ["./opaque-server-binary"]
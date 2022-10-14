FROM golang:1.17-alpine

WORKDIR /app

COPY ./opaque-client/go.mod ./
COPY ./opaque-client/go.sum ./

RUN go mod download

COPY ./opaque-client /app
COPY ./opaque-service-protos /app

RUN cd ./opaque-client
RUN go build -o /opaque-client

CMD [ "/bin/bash" ]
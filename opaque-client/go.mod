module github.com/signed-long/opaque-over-grpc/opaque-client

go 1.17

replace github.com/signed-long/opaque-over-grpc/opaque-service-protos v0.0.0 => ../opaque-service-protos

require (
	github.com/cretz/gopaque v0.1.0
	github.com/signed-long/opaque-over-grpc/opaque-service-protos v0.0.0
	github.com/sirupsen/logrus v1.9.0
	github.com/tyler-smith/go-bip39 v1.1.0
	golang.org/x/crypto v0.0.0-20220214200702-86341886e292
	golang.org/x/term v0.0.0-20220722155259-a9ba230a4035
	google.golang.org/grpc v1.48.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	go.dedis.ch/fixbuf v1.0.3 // indirect
	go.dedis.ch/kyber/v3 v3.0.12 // indirect
	golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

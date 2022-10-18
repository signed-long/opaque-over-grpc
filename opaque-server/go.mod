module github.com/signed-long/opaque-over-grpc/opaque-server

go 1.17

replace github.com/signed-long/opaque-over-grpc/opaque-service-protos v0.0.0 => ../opaque-service-protos

require (
	github.com/cretz/gopaque v0.1.0
	github.com/google/uuid v1.1.2
	github.com/signed-long/opaque-over-grpc/opaque-service-protos v0.0.0
	github.com/sirupsen/logrus v1.9.0
	google.golang.org/grpc v1.48.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/stretchr/testify v1.7.2 // indirect
	go.dedis.ch/fixbuf v1.0.3 // indirect
	go.dedis.ch/kyber/v3 v3.0.12 // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 // indirect
	golang.org/x/net v0.0.0-20201021035429-f5854403a974 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

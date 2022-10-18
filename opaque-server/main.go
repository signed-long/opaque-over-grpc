package main

import (
	"net"
	"os"

	logrus "github.com/sirupsen/logrus"

	authServiceImplementation "github.com/signed-long/opaque-over-grpc/opaque-server/svc_auth"
	protos "github.com/signed-long/opaque-over-grpc/opaque-service-protos/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := logrus.New()
	log.Out = os.Stdout
	log.SetLevel(logrus.DebugLevel)
	// log.SetReportCaller(true)

	grpcServer := grpc.NewServer()
	authServiceServer := authServiceImplementation.NewAuthServiceServer(*log)

	protos.RegisterOpaqueAuthServiceServer(grpcServer, authServiceServer)

	reflection.Register(grpcServer) // TODO: turn this off with config DEBUG

	l, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatal("Unable to listen error: ", err)
	}

	grpcServer.Serve(l)
}

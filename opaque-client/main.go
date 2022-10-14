package main

import (
	"flag"
	"os"

	logrus "github.com/sirupsen/logrus"

	cli "github.com/signed-long/opaque-over-grpc/opaque-client/cli"
	authClientImplementaion "github.com/signed-long/opaque-over-grpc/opaque-client/client_auth"
	"github.com/signed-long/opaque-over-grpc/opaque-service-protos/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	log := logrus.New()
	log.Out = os.Stdout
	log.SetLevel(logrus.DebugLevel)
	// log.SetReportCaller(true)

	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	grpcClinet := protos.NewOpaqueAuthServiceClient(conn)
	authClient := authClientImplementaion.NewAuthServiceClient(grpcClinet, *log)

	registerCommand := flag.NewFlagSet("reg", flag.ExitOnError)

	if len(os.Args) < 2 {
		cli.PrintCommandSummary()
		os.Exit(1)
	}

	// look at the 2nd argument's value
	switch os.Args[1] {
	case "reg":
		cli.HandleReg(registerCommand, *authClient)
	default:
		cli.PrintCommandSummary()
	}
}

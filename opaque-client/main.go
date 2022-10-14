package main

import (
	"flag"
	"os"

	logrus "github.com/sirupsen/logrus"

	"github.com/signed-long/opaque-over-grpc/opaque-client/cli"
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
	regHelp := registerCommand.Bool("help", false, "Get more information about this command.")

	addCommand := flag.NewFlagSet("add", flag.ExitOnError)
	addHIBPHashCheck := addCommand.Bool("hipb", true, "Check candidate assword against the HIPB database of leaked passwords.")
	addLogin := addCommand.String("login", "", "Add a login assocaiated with the new password eg. email or username.")
	addHelp := addCommand.Bool("help", false, "Get more information about this command.")

	copyCommand := flag.NewFlagSet("copy", flag.ExitOnError)
	copyHelp := copyCommand.Bool("help", false, "Get more information about this command.")

	printCommand := flag.NewFlagSet("print", flag.ExitOnError)
	printHelp := printCommand.Bool("help", false, "Get more information about this command.")

	deleteCommand := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteHelp := deleteCommand.Bool("help", false, "Get more information about this command.")

	generateCommand := flag.NewFlagSet("gen", flag.ExitOnError)
	genHelp := generateCommand.Bool("help", false, "Get more information about this command.")

	if len(os.Args) < 2 {
		cli.PrintCommandSummary()
		os.Exit(1)
	}

	// look at the 2nd argument's value
	switch os.Args[1] {
	case "reg":
		cli.HandleReg(registerCommand, regHelp, *authClient)
	case "add":
		cli.HandleAdd(addCommand, addHIBPHashCheck, addLogin, addHelp)
	case "copy":
		cli.HandleCopy(copyCommand, copyHelp)
	case "print":
		cli.HandlePrint(printCommand, printHelp)
	case "delete":
		cli.HandleDelete(deleteCommand, deleteHelp)
	case "gen":
		cli.HandleGen(generateCommand, genHelp)
	default:
		cli.PrintCommandSummary()
	}
}

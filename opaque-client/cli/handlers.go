package cli

import (
	"flag"
	"fmt"
	"os"

	opaque "github.com/signed-long/opaque-over-grpc/opaque-client/client_auth"
)

func HandleReg(regCommand *flag.FlagSet, help *bool, authServiceClient opaque.AuthServiceClient) {
	regCommand.Parse(os.Args[2:])
	if *help {
		printAndQuit("This is a help message for the reg command")
	}

	fmt.Println("Enter a username and password to register")

	msg := "email: "
	email, err := sensitiveInput(msg)
	if err != nil {
		fmt.Println("Failed to read email")
		os.Exit(1)
	}
	msg = "password: "
	password, err := sensitiveInput(msg)
	if err != nil {
		fmt.Println("Failed to read password")
		os.Exit(1)
	}

	err = authServiceClient.RegisterFlow(email, password)
	if err != nil {
		// s, _ := status.FromError(err)
		// if s.Code() == codes.AlreadyExists {
		// 	fmt.Println("Registration failed: An account has already been registered for the email provided")
		// 	os.Exit(1)
		// }
		fmt.Println("Registration failed: ", err)
		os.Exit(1)
	}

	fmt.Println("Registration successful")
}

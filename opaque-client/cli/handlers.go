package cli

import (
	"flag"
	"fmt"
	"os"

	opaque "github.com/signed-long/opaque-over-grpc/opaque-client/client_auth"
)

func HandleReg(regCommand *flag.FlagSet, authServiceClient opaque.AuthServiceClient) {
	regCommand.Parse(os.Args[2:])

	fmt.Println("Enter a username and password to register")

	var username, password string
	login(&username, &password)

	err := authServiceClient.RegisterFlow(username, password)
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

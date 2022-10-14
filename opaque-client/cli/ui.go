package cli

import (
	"fmt"
	"os"
	"syscall"

	"golang.org/x/term"
)

func PrintCommandSummary() {
	fmt.Println("commands:\n")
	fmt.Printf("register\tregister with the server\n")
	fmt.Printf("authenticate\tauthenticate with the server, a session token will be issued\n")
	fmt.Println()
}

func printAndQuit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func sensitiveInput(msg string) (string, error) {
	fmt.Println(msg)
	byteInput, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}

	input := string(byteInput)
	return input, nil
}

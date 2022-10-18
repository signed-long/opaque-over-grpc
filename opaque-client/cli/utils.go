package cli

import (
	"fmt"
	"os"
	"syscall"

	"golang.org/x/term"
)

func isAllowedChar(char rune) bool {
	isAlpha := ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z')
	isNum := ('0' <= char && char <= '9')
	isUnderscore := (char == '-')
	return (isAlpha || isNum || isUnderscore)
}

func login(username *string, password *string) {
	var err error
	msg := "username: "
	*username, err = sensitiveInput(msg)
	if err != nil {
		fmt.Println("Failed to read email")
		os.Exit(1)
	}
	msg = "password: "
	*password, err = sensitiveInput(msg)
	if err != nil {
		fmt.Println("Failed to read password")
		os.Exit(1)
	}
}

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

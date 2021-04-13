package runner

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
	"syscall"
)

func GetPasswordEntryFromUser() (string, string, []byte) {
	input := bufio.NewReader(os.Stdin)

	fmt.Printf("Entry name: ")
	entryName, _ := input.ReadString('\n')

	fmt.Printf("Email/Username: ")
	userName, _ := input.ReadString('\n')

	fmt.Printf("Password: ")
	password, _ := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Printf("\n")

	return strings.TrimSuffix(entryName, "\n"), strings.TrimSuffix(userName, "\n"), password
}

func GetEntryNameFromUser() string {
	input := bufio.NewReader(os.Stdin)

	fmt.Printf("Entry name: ")
	entryName, _ := input.ReadString('\n')

	return strings.TrimSuffix(entryName, "\n")
}
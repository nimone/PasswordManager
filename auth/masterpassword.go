package auth

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"syscall"
)

func CreateMasterPassword() []byte {
	var masterPassword []byte

	fmt.Println("Create a master password (make sure it's a strong one i.e. min 8 characters)")

	for {
		fmt.Printf("Master Password: ")
		masterPassword, _ = terminal.ReadPassword(int(syscall.Stdin))
		fmt.Printf("\n")

		if len(string(masterPassword)) < 8 {
			fmt.Println("Minimum password length should be 8 characters, try again")
			continue
		}

		fmt.Printf("Confirm Master Password: ")
		masterPasswordConfirm, _ := terminal.ReadPassword(int(syscall.Stdin))
		fmt.Printf("\n")

		if string(masterPassword) != string(masterPasswordConfirm) {
			fmt.Println("Passwords doesn't match, try again")
			continue
		}

		fmt.Println("Master password created successfully")
		break
	}
	return masterPassword
}

func GetMasterPassword() []byte {
	var masterPassword []byte

	fmt.Printf("Master Password: ")
	masterPassword, _ = terminal.ReadPassword(int(syscall.Stdin))
	fmt.Printf("\n")

	return masterPassword
}
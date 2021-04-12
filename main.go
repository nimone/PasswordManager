package main

import (
	"bufio"
	"fmt"
	"github.com/nimone/PasswordManager/crypto"
	"github.com/nimone/PasswordManager/db"
	"github.com/nimone/PasswordManager/helper"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
	"syscall"
)

const dbPath = "./test.db"

func createMasterPassword() []byte {
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

		fmt.Printf("Confirm: ")
		masterPasswordConfirm, _ := terminal.ReadPassword(int(syscall.Stdin))
		fmt.Printf("\n")

		if string(masterPassword) != string(masterPasswordConfirm) {
			fmt.Println("Passwords doesn't match, try again")
			continue
		}
		break
	}
	return masterPassword
}

func getMasterPassword() []byte {
	var masterPassword []byte

	fmt.Printf("Master Password: ")
	masterPassword, _ = terminal.ReadPassword(int(syscall.Stdin))
	fmt.Printf("\n")

	return masterPassword
}

func main() {
	fmt.Println("Go Password Manager")

	var masterPassword []byte
	firstRun := false

	// if the database is does not exist (first run)
	// prompt the use to create a master password
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		masterPassword = createMasterPassword()
		firstRun = true
	}

	passwordDB, err := db.Init(dbPath)
	helper.HandleError(err)

	// create the masterpassword entry (first entry)
	if firstRun {
		masterPasswordHash, _ := crypto.HashPassword(masterPassword)

		tx := passwordDB.Create(&db.PasswordEntry{
			Password: string(masterPasswordHash),
		})
		helper.HandleError(tx.Error)
		err = db.GenerateDummyEntries(passwordDB)
		helper.HandleError(err)

		// else validate the user
	} else {
		var masterPasswordEntry db.PasswordEntry
		passwordDB.First(&masterPasswordEntry)

		err := crypto.CheckPasswordHash(
			[]byte(masterPasswordEntry.Password),
			getMasterPassword(),
		)
		if err != nil {
			fmt.Println("Master password you provided is invalid")
			os.Exit(1)
		} else {
			fmt.Println("Autheticated Successfully")
		}
	}

	var mainOpt int
	fmt.Printf("What do you want to do?\n1. Store a password\n2. Retrieve a password\n> ")
	fmt.Scanf("%d\n", &mainOpt)

	input := bufio.NewReader(os.Stdin)

	switch mainOpt {
	case 1:
		fmt.Printf("Entry name: ")
		entryName, _ := input.ReadString('\n')

		fmt.Printf("Email/username: ")
		userName, _ := input.ReadString('\n')

		fmt.Printf("Password: ")
		password, _ := input.ReadString('\n')

		tx := passwordDB.Create(&db.PasswordEntry{
			EntryName: strings.TrimSuffix(entryName, "\n"),
			UserName:  strings.TrimSuffix(userName, "\n"),
			Password:  strings.TrimSuffix(password, "\n"),
		})
		if !helper.HandleError(tx.Error) {
			fmt.Println("The password is saved successfully")
		}

		break

	case 2:
		var passwordEntries []db.PasswordEntry

		fmt.Printf("Retrive the password for: ")
		entryName, _ := input.ReadString('\n')
		entryName = strings.TrimSuffix(entryName, "\n")

		tx := passwordDB.Where(
			"entry_name LIKE ?",
			"%"+entryName+"%",
		).Find(&passwordEntries)

		if !helper.HandleError(tx.Error) {
			if len(passwordEntries) == 0 {
				fmt.Printf("There's nothing for '%s'\n", entryName)

			} else {
				fmt.Println("EntryName\t Username\t Password")
				for _, e := range passwordEntries {
					fmt.Printf("%s\t %s\t %s\n", e.EntryName, e.UserName, e.Password)
				}
			}
		}
		break

	default:
		fmt.Println("Don't know what you're talking about")
	}
}
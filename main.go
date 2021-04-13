package main

import (
	"bufio"
	"fmt"
	"github.com/nimone/PasswordManager/auth"
	"github.com/nimone/PasswordManager/crypto"
	"github.com/nimone/PasswordManager/db"
	"github.com/nimone/PasswordManager/helper"
	// "golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
	// "syscall"
)

const dbPath = "./test.db"

func main() {
	fmt.Println("Go Password Manager")

	var masterPassword []byte
	firstRun := false

	// if the database does not exist (first run)
	// prompt the use to create a master password
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		masterPassword = auth.CreateMasterPassword()
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

		// else authenticate the user
	} else {
		var masterPasswordEntry db.PasswordEntry
		masterPassword = auth.GetMasterPassword()

		passwordDB.First(&masterPasswordEntry)

		UserAuthenticated := auth.AuthenticateUser(
			[]byte(masterPasswordEntry.Password),
		)

		if UserAuthenticated {
			fmt.Println("Authetication Successful")
		} else {
			fmt.Println("Authentication Failed: Master password is invalid")
			os.Exit(1)
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
package main

import (
	"bufio"
	"fmt"
	"github.com/nimone/PasswordManager/db"
	"github.com/nimone/PasswordManager/helper"
	"os"
	"strings"
)

func main() {
	fmt.Println("Go Password Manager")

	passwordDB, err := db.Init("./test.db")
	helper.HandleError(err)

	err = db.GenerateDummyEntries(passwordDB)
	helper.HandleError(err)

	var mainOpt int
	fmt.Printf("What do you want to do?\n1. Store a password\n2. Retrieve a password\n> ")
	fmt.Scanf("%d\n", &mainOpt)

	input := bufio.NewReader(os.Stdin)

	switch mainOpt {
	case 1:
		fmt.Printf("The password is for: ")
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
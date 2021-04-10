package main

import (
	"fmt"
	"github.com/nimone/PasswordManager/db"
	"github.com/nimone/PasswordManager/helper"
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

	switch mainOpt {
	case 1:
		passwordEntry := db.PasswordEntry{}

		fmt.Printf("The password is for: ")
		fmt.Scanf("%s\n", &passwordEntry.EntryName)

		fmt.Printf("Email/username: ")
		fmt.Scanf("%s\n", &passwordEntry.UserName)

		fmt.Printf("Password: ")
		fmt.Scanf("%s\n", &passwordEntry.Password)

		tx := passwordDB.Create(&passwordEntry)
		if !helper.HandleError(tx.Error) {
			fmt.Println("The password is saved successfully")
		}

		break

	case 2:
		var entryName string
		var passwordEntries []db.PasswordEntry

		fmt.Printf("Retrive the password for: ")
		fmt.Scanf("%s\n", &entryName)

		tx := passwordDB.Where(
			"entry_name LIKE ?", "%"+entryName+"%",
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
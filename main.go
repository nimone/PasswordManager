package main

import (
	"fmt"
	"github.com/nimone/PasswordManager/db"
)

func main() {
	fmt.Println("Go Password Manager")

	passwordDB := db.PasswordDB{}
	passwordDB.Init("./test.db")

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

		err := passwordDB.Store(passwordEntry)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("The password is saved successfully")
		}

		break

	case 2:
		var entryName string
		fmt.Printf("Retrive the password for: ")
		fmt.Scanf("%s\n", &entryName)

		entries, err := passwordDB.Get(entryName)
		if err != nil {
			fmt.Println(err)

		} else {
			if len(entries) == 0 {
				fmt.Printf("There's no entry for '%s'", entryName)

			} else {
				fmt.Println("EntryName\t Username\t Password")
				for _, e := range entries {
					fmt.Printf("%s\t %s\t %s\n", e.EntryName, e.UserName, e.Password)
				}
			}
		}
		break

	default:
		fmt.Println("Don't know what you're talking about")
	}
}
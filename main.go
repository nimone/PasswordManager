package main

import (
	"encoding/base64"
	"fmt"
	"github.com/nimone/PasswordManager/auth"
	"github.com/nimone/PasswordManager/crypto"
	"github.com/nimone/PasswordManager/db"
	"github.com/nimone/PasswordManager/helper"
	"github.com/nimone/PasswordManager/runner"
	"os"
)

const dbPath = "./test.db"

// AES-256 encryption | using base64 encoding for database storage purpose
func EncryptPassword(password []byte, key *[]byte) (string, error) {
	passwordCipher, err := crypto.Encrypt(password, key)

	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(passwordCipher), nil
}

func DecryptPassword(b64PasswordCipher string, key *[]byte) (string, error) {
	passwordCipher, _ := base64.StdEncoding.DecodeString(b64PasswordCipher)
	passwordInHex, err := crypto.Decrypt(passwordCipher, key)

	if err != nil {
		return "", err
	}
	return string(passwordInHex), nil
}

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

	// On first run create the masterpassword entry (first entry)
	if firstRun {
		masterPasswordHash, _ := crypto.HashPassword(masterPassword)

		tx := passwordDB.Create(&db.PasswordEntry{
			Password: string(masterPasswordHash),
		})
		helper.HandleError(tx.Error)
		// err = db.GenerateDummyEntries(passwordDB) // currently not encrypting dummy entries
		// helper.HandleError(err)

		// else authenticate the user
	} else {
		var masterPasswordEntry db.PasswordEntry
		masterPassword = auth.GetMasterPassword()

		passwordDB.First(&masterPasswordEntry)

		UserAuthenticated := auth.AuthenticateUser(
			[]byte(masterPasswordEntry.Password), // master password hash
			masterPassword,
		)

		if UserAuthenticated {
			fmt.Println("Authetication Successful")
		} else {
			fmt.Println("Authentication Failed: Master password is invalid")
			os.Exit(1)
		}
	}

	// Maybe a bad idea to have hardcoded salt
	const salt = "somesaltysalt"
	// Generate a 256-bit key from masterpassword to encrypt/decrypt
	key, err := crypto.GenerateKey(masterPassword, []byte(salt))

	if helper.HandleError(err) {
		fmt.Println("Failed to generate the encryption key")
		os.Exit(2)
	}

	// ===================================
	// Program Functionality starts here
	// ===================================

	var mainOpt int
	fmt.Printf("What do you want to do?\n1. Store a password\n2. Retrieve a password\n> ")
	fmt.Scanf("%d\n", &mainOpt)

	switch mainOpt {
	case 1:
		entryName, userName, password := runner.GetPasswordEntryFromUser()

		b64PasswordCipher, err := EncryptPassword(password, &key)
		if helper.HandleError(err) {
			fmt.Println("Failed to encrypt the password")
			os.Exit(2)
		}

		tx := passwordDB.Create(&db.PasswordEntry{
			EntryName: entryName,
			UserName:  userName,
			Password:  b64PasswordCipher,
		})
		if !helper.HandleError(tx.Error) {
			fmt.Println("The password is saved successfully")
		}

		break

		// Refactor needed
	case 2:
		var passwordEntries []db.PasswordEntry
		entryName := runner.GetEntryNameFromUser()

		tx := passwordDB.Where(
			"entry_name LIKE ?", "%"+entryName+"%",
		).Find(&passwordEntries)

		if helper.HandleError(tx.Error) {
			fmt.Printf("Couldn't query the data for '%s'", entryName)
			os.Exit(3)
		}

		if len(passwordEntries) == 0 {
			fmt.Printf("Nothing found for '%s'\n", entryName)

		} else if len(passwordEntries) == 1 {
			password, err := DecryptPassword(passwordEntries[0].Password, &key)
			if helper.HandleError(err) {
				fmt.Println("Failed to decrypt the password")
				os.Exit(2)
			}
			fmt.Println("Email/Username:", passwordEntries[0].UserName)
			fmt.Println("Password:", password)

		} else {
			var entryID uint
			var passwordEntry db.PasswordEntry

			fmt.Printf("Found %d entries for '%s'\n", len(passwordEntries), entryName)
			fmt.Println("EntryID\t EntryName\t Username\t")

			for _, e := range passwordEntries {
				fmt.Printf(
					"%d\t %s\t %s\n",
					e.ID, e.EntryName, e.UserName,
				)
			}
			fmt.Printf("\nEntryID: ")
			fmt.Scanf("%d", &entryID)

			tx := passwordDB.First(&passwordEntry, entryID)
			if helper.HandleError(tx.Error) {
				fmt.Println("Couldn't find the EntryID:", entryID)

			} else {
				password, err := DecryptPassword(passwordEntry.Password, &key)
				if helper.HandleError(err) {
					fmt.Println("Failed to decrypt the password")
					os.Exit(2)
				}
				fmt.Println("Email/Username:", passwordEntry.UserName)
				fmt.Println("Password:", password)
			}
		}

		break

	default:
		fmt.Println("Don't know what you're talking about")
	}
}
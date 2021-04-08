package main

import (
	"fmt"
	"github.com/nimone/PasswordManager/db"
)

func main() {
	fmt.Println("Go Password Manager")
	db.Init("./test.db")
}
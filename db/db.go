package db

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nimone/PasswordManager/helper"
)

func Init(dbpath string) error {
	db, err := sql.Open("sqlite3", dbpath)

	if helper.HandleError(err) {
		return errors.New("Cannot open the path specified")
	}

	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS Sources (id INTEGER PRIMARY KEY, name TEXT)")
	helper.HandleError(err)
	statement.Exec()

	statement, err = db.Prepare("CREATE TABLE IF NOT EXISTS Passwords (id INTEGER PRIMARY KEY, source integer FORIEGN KEY, username TEXT, email TEXT, password TEXT)")
	helper.HandleError(err)
	statement.Exec()

	return nil
}
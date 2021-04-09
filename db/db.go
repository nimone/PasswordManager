package db

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nimone/PasswordManager/helper"
)

type PasswordEntry struct {
	ID        int
	EntryName string
	UserName  string
	Password  string
}

type PasswordDB struct {
	db *sql.DB
}

func (pass *PasswordDB) Init(dbpath string) error {
	db, err := sql.Open("sqlite3", dbpath)

	if helper.HandleError(err) {
		return errors.New("Cannot open the path specified")
	}
	pass.db = db

	// statement, err := pass.db.Prepare(`
	// 	CREATE TABLE IF NOT EXISTS
	// 		Sources (
	// 			id INTEGER PRIMARY KEY,
	// 			name TEXT
	// 		)
	//;`)
	// if helper.HandleError(err) {
	// 	return errors.New("Cannot initilize Sources table")
	// statement.Exec()

	statement, err := pass.db.Prepare(`
		CREATE TABLE IF NOT EXISTS 
			Passwords (
				id INTEGER PRIMARY KEY, 
				entryname TEXT, 
				username TEXT, 
				password TEXT
			)
	;`)
	if helper.HandleError(err) {
		return errors.New("Cannot initilize Passwords table")
	}
	statement.Exec()

	return nil
}

func (pass *PasswordDB) Store(entry PasswordEntry) error {
	statement, err := pass.db.Prepare(`
		INSERT INTO 
			Passwords (entryname, username, password) 
			VALUES (?,?,?)
	;`)
	if helper.HandleError(err) {
		return errors.New("Cannot create a password entry")
	}
	statement.Exec(entry.EntryName, entry.UserName, entry.Password)

	return nil
}

func (pass *PasswordDB) Get(entryName string) ([]*PasswordEntry, error) {
	rows, err := pass.db.Query(`
		SELECT * FROM Passwords WHERE entryname LIKE ?;`,
		"%"+entryName+"%",
	)
	if helper.HandleError(err) {
		return nil, errors.New("Cannot retrive the password entry")
	}
	defer rows.Close()

	var result []*PasswordEntry

	for rows.Next() {
		entry := &PasswordEntry{}
		rows.Scan(&entry.ID, &entry.EntryName, &entry.UserName, &entry.Password)
		result = append(result, entry)
	}
	return result, nil
}
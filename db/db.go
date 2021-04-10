package db

import (
	"errors"
	"github.com/nimone/PasswordManager/helper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type PasswordEntry struct {
	ID        uint `gorm:"primaryKey"`
	EntryName string
	UserName  string
	Password  string
}

func Init(dbpath string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbpath), &gorm.Config{})

	if helper.HandleError(err) {
		return nil, errors.New("failed to connect database")
	}
	err = db.AutoMigrate(&PasswordEntry{})
	if helper.HandleError(err) {
		return nil, errors.New("Cannot migrating the schema")
	}

	return db, nil
}

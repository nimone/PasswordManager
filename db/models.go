package db

type PasswordEntry struct {
	ID        uint `gorm:"primaryKey"`
	EntryName string
	UserName  string
	Password  string
}
package db

import (
	"gorm.io/gorm"
	"math/rand"
)

var (
	usernames = []string{"fishdaily", "performerpoppy", "gulliblemembrane", "suspendadaptable", "almostprefer",
		"penitentgolfing", "copemit", "legwarmersworkman", "firstvulcan", "cranberriescobra",
		"fearlessoval", "dweebperky", "inquirytasteful", "dirtwhich", "halftimesternway",
		"anybodycroissant", "bellowindustry", "mileomelette", "takejaialai", "lessontendency",
	}
	passwords = []string{"RAd7RHYU", "PpaxWXhs", "uuB5BQyV", "RYexFz3R", "kfJ2GDXh",
		"Dv5cmn7Y", "QELbndxN", "NQGH2fqU", "KzfjVANb", "dRCKGcrG",
		"yWfjr38K", "5DkYAyJc", "CSUjVpyP", "Fe6hcvAJ", "7QdkFtsK",
		"fFSDY8Vs", "vQvb5Z7U", "YuMjrRS2", "geaVaxu5", "jD5hkGsB",
	}
	entryNames = []string{"google.com", "facebook.com", "system login", "twitter.com",
		"AWS", "Digital Ocean", "github.com", "heroku", "crypto.com", "notion",
	}
)

func GenerateDummyEntries(passDB *gorm.DB) error {
	for _, entryName := range entryNames {
		tx := passDB.Create(&PasswordEntry{
			EntryName: entryName,
			UserName:  usernames[rand.Intn(len(usernames))],
			Password:  passwords[rand.Intn(len(passwords))],
		})

		if tx.Error != nil {
			return tx.Error
		}
	}
	return nil
}
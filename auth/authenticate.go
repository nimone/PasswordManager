package auth

import (
	"github.com/nimone/PasswordManager/crypto"
)

func AuthenticateUser(masterPasswordHash, masterPassword []byte) bool {
	err := crypto.CheckPasswordHash(
		masterPasswordHash,
		masterPassword,
	)
	return err == nil
}
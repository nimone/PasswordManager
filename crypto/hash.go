package crypto

import (
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"
)

func HashPassword(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
}

// securely compares a bcrypt hashed password with an unhashed password
// returns nil on success and an error on failure
func CheckPasswordHash(hash, password []byte) error {
	return bcrypt.CompareHashAndPassword(hash, password)
}

// Generates a 256-bit key from a variable length password
// for aes encryption and decryption
func GenerateKey(password, salt []byte) ([]byte, error) {
	const (
		iterations = 16384
		memorycost = 8
		cpucost    = 1
		keylength  = 32
	)
	return scrypt.Key(password, salt, iterations, memorycost, cpucost, keylength)
}

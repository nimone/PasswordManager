package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

// encrypts data using 256-bit AES-GCM and provides a check that it hasn't been altered.
func Encrypt(plaintext []byte, key *[]byte) (ciphertext []byte, err error) {
	block, err := aes.NewCipher(*key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

// decrypts data using 256-bit AES-GCM and provides a check that it hasn't been altered.
func Decrypt(ciphertext []byte, key *[]byte) (plaintext []byte, err error) {
	block, err := aes.NewCipher(*key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < gcm.NonceSize() {
		return nil, errors.New("malformed ciphertext")
	}

	return gcm.Open(nil,
		ciphertext[:gcm.NonceSize()],
		ciphertext[gcm.NonceSize():],
		nil,
	)
}

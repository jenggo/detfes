package pkg

import (
	"golang.org/x/crypto/bcrypt"
)

func EncryptAPIKEY(input string) string {
	i := []byte(input)
	b, err := bcrypt.GenerateFromPassword(i, 4)
	if err != nil {
		return ""
	}

	return string(b)
}

func ValidAPIKey(hashed string, plaintext string) bool {
	if cacheInvalidKey(hashed) {
		return false
	}

	p := []byte(plaintext)
	h := []byte(hashed)
	if err := bcrypt.CompareHashAndPassword(h, p); err != nil {
		return false
	}

	return true
}

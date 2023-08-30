package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(raw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), 10)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ComparePassword(hash, raw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(raw))
}

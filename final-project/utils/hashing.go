package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(p string) string {
	salt := 16
	password := []byte(p)
	hash, err := bcrypt.GenerateFromPassword(password, salt)

	if err != nil {
		log.Fatal("Failed to hash password", err)
	}

	return string(hash)
}

func ComparePassword(h, p []byte) bool {
	hash, pass := []byte(h), []byte(p)

	err := bcrypt.CompareHashAndPassword(hash, pass)

	return err == nil
}

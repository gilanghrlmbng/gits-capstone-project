package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password, id string) string {
	pass := []byte(fmt.Sprintf("%s:%s", password, id))

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

func CheckPassword(password, id, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(fmt.Sprintf("%s:%s", password, id)))
	return err == nil
}

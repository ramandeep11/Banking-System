package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string ,error) {
	hasedpass, err := bcrypt.GenerateFromPassword([]byte(password) , bcrypt.DefaultCost)

	if err!= nil {
		return "", fmt.Errorf("failed to hash password")
	}

	return string(hasedpass),nil
}

func CheckPassword(password string , hashedpass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedpass), []byte(password))
}
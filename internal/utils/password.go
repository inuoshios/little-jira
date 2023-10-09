package utils

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func ComparePasswords(hashedPassword, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		log.Println(err.Error())
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return ErrMismatchedHashAndPassword
		default:
			return ErrIncorrectPassword
		}
	}

	return nil
}

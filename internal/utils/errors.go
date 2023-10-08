package utils

import "errors"

var (
	ErrInvalidToken              = errors.New("unexpected signing method")
	ErrSignatureInvalid          = errors.New("signature is invalid")
	ErrTokenExpired              = errors.New("token is expired")
	ErrTokenNotValidYet          = errors.New("token is not valid yet")
	ErrMismatchedHashAndPassword = errors.New("password mismatch, please try again")
)

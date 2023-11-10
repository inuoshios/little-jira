package utils

import "errors"

var (
	ErrInvalidToken              = errors.New("unexpected signing method")
	ErrSignatureInvalid          = errors.New("signature is invalid")
	ErrTokenExpired              = errors.New("token is expired")
	ErrTokenNotValidYet          = errors.New("token is not valid yet")
	ErrMismatchedHashAndPassword = errors.New("password mismatch, please try again")
	ErrSqlNoRowsUser             = errors.New("user does not exist")
	ErrIncorrectPassword         = errors.New("invalid credentials")
	ErrAuthHeader                = errors.New("authorization header not provided")
	ErrInvalidAuthHeader         = errors.New("invalid authorization header format")
	ErrUnsupportedAuthType       = errors.New("unsupported authorization type")
	ErrAuthorizationError        = errors.New(("an error occured, please try again later"))
)

package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

func NewClaims(userId, username, email string, duration time.Duration) (*Claims, error) {
	id, _ := uuid.NewRandom()
	claims := &Claims{
		UserID:   userId,
		Username: username,
		Email:    email,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        id.String(),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}

	return claims, nil
}

// GenerateToken generates a new JWT token
func GenerateToken(userId, username, email string, duration time.Duration) (string, error) {
	payload, err := NewClaims(userId, username, email, duration)
	if err != nil {
		return "", fmt.Errorf("jwt error: %v", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	signedStr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", fmt.Errorf("error signing jwt token: %v", err)
	}

	return signedStr, nil
}

// VerifyToken verifies the JWT token that was generated
func VerifyToken(tokenString string) (*Claims, error) {
	keyFunc := func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	}
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, keyFunc)
	switch {
	case errors.Is(err, jwt.ErrSignatureInvalid):
		return nil, ErrSignatureInvalid
	case errors.Is(err, jwt.ErrTokenExpired):
		return nil, ErrTokenExpired
	case errors.Is(err, jwt.ErrTokenNotValidYet):
		return nil, ErrTokenNotValidYet
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}

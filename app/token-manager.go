package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type TokenManager struct {
	secretKey []byte
}

func (tokenManager *TokenManager) createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(tokenManager.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (tokenManager *TokenManager) validate(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return tokenManager.secretKey, nil
	})

	if err != nil {
		return false, err
	}

	if token.Valid {
		return true, nil
	}

	return false, fmt.Errorf("invalid token")
}

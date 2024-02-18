package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tinycloudtv/authn-service/app/internal/models"
	"github.com/tinycloudtv/authn-service/app/internal/repositories"
	"time"
)

type TokenManager struct {
	secretKey []byte
}

func (tokenManager *TokenManager) createToken(user models.User) (string, error) {
	expiration := time.Now().Add(time.Hour * 24)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": user.Email,
			"exp":      expiration.Unix(),
		})

	tokenString, err := token.SignedString(tokenManager.secretKey)
	if err != nil {
		return "", err
	}

	userTokenRepo := repositories.UserTokensRepository{}
	userTokenRepo.Create(user, tokenString, expiration.UTC())

	return tokenString, nil
}

func (tokenManager *TokenManager) validate(tokenString string) (bool, error) {
	userTokenRepo := repositories.UserTokensRepository{}
	userToken, dberr := userTokenRepo.Get(tokenString)

	if dberr != nil {
		return false, dberr
	}

	if userToken.Expired {
		return false, nil
	}

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

package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)


type JWTClamis struct {
	UserId string `json:"userId"`
	jwt.RegisteredClaims
}

func CreateToken(userId string, secret string) (string, error) {
	claims := JWTClamis{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
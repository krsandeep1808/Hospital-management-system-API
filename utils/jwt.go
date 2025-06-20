package utils

import (
	"clinic-portal/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateToken(role string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"role": role,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, _ := token.SignedString([]byte(config.EnvJWTSecret()))
	return tokenString
}

func ValidateToken(tokenString string, expectedRole string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.EnvJWTSecret()), nil
	})

	if err != nil || token == nil {
		return false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["role"] == expectedRole
	}

	return false
}

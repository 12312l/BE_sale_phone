package utils

import (
	"time"

	"daoduy.com/hoc-golang/config"
	"github.com/golang-jwt/jwt/v5"
)

// Tạo JWT token
func GenerateToken(userID uint, username string, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"username": username,
		"scope":   role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // hết hạn sau 24h
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.JWTSecret)
}

// Xác minh JWT token
func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return config.JWTSecret, nil
	})
}

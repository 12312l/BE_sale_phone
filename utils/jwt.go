package utils

import (
	"time"

	"github.com/google/uuid"
	"daoduy.com/hoc-golang/config"
	"daoduy.com/hoc-golang/models/response"
	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken sinh JWT cho người dùng
func GenerateToken(userResponse response.UserResponse) (string, error) {
	claims := jwt.MapClaims{
		"jti":      uuid.NewString(),
		"user_id":  userResponse.ID,
		"username": userResponse.Username,
		"scope":    userResponse.Role,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.JWTSecret)
}

// ValidateToken xác minh token JWT
func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return config.JWTSecret, nil
	})
}

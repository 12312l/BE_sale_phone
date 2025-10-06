package middleware

import (
	"net/http"
	"strings"

	"daoduy.com/hoc-golang/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(requiredRole ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Thiếu token"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := utils.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token không hợp lệ"})
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		role := claims["scope"].(string) // 👈 Lấy role từ claim

		// Nếu middleware yêu cầu role cụ thể
		if len(requiredRole) > 0 && role != requiredRole[0] {
			c.JSON(http.StatusForbidden, gin.H{"error": "Không có quyền truy cập"})
			c.Abort()
			return
		}

		c.Set("user_id", claims["user_id"])
		c.Set("role", role)
		c.Next()
	}
}

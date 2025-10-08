package middleware

import (
	"net/http"
	"strings"

	"daoduy.com/hoc-golang/config"
	"daoduy.com/hoc-golang/exception"
	"daoduy.com/hoc-golang/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := strings.TrimSpace(c.GetHeader("Authorization"))
		if authHeader == "" {
			// c.JSON(http.StatusUnauthorized, gin.H{"error": "Thiếu token"})
			c.Error(exception.NewAppException(exception.Unauthenticated))
			c.Abort()
			return
		}

		// Hỗ trợ cả "Bearer <token>" và "<token>"
		var tokenString string
		if strings.HasPrefix(authHeader, "Bearer ") {
			tokenString = strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
		} else {
			tokenString = authHeader
		}

		token, err := utils.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			c.Error(exception.NewAppException(exception.Unauthenticated))
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Không đọc được claims"})
			c.Abort()
			return
		}

		jti, _ := claims["jti"].(string)

		// ✅ Kiểm tra jti có bị vô hiệu hóa không
		var count int64
		config.DB.Table("invalidated_tokens").Where("id = ?", jti).Count(&count)
		if count > 0 {
			c.Error(exception.NewAppException(exception.Unauthenticated))
			c.Abort()
			return
		}

		role, _ := claims["scope"].(string)

		if len(requiredRoles) > 0 {
			hasRole := false
			for _, r := range requiredRoles {
				if role == r {
					hasRole = true
					break
				}
			}
			if !hasRole {
				c.Error(exception.NewAppException(exception.Unauthorized))
				c.Abort()
				return
			}
		}
		userIDFloat, ok := claims["user_id"].(float64)
		if ok {
			c.Set("user_id", uint(userIDFloat))
		}
		c.Set("username", claims["username"])
		c.Set("role", role)

		c.Next()
	}
}

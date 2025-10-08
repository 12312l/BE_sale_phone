package services

import (
	"errors"
	"os"
	"time"

	"daoduy.com/hoc-golang/config"
	"daoduy.com/hoc-golang/exception"
	"daoduy.com/hoc-golang/models"
	request "daoduy.com/hoc-golang/models/request"
	"daoduy.com/hoc-golang/models/response"
	"daoduy.com/hoc-golang/utils"
	"github.com/golang-jwt/jwt/v5"
)

func Login(req request.LoginRequest) (string, error) {
	var user models.User
	if err := config.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		return "", errors.New("Người dùng không tồn tại")
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return "", errors.New("Mật khẩu không chính xác")
	}

	// 👇 Chuẩn bị UserResponse để truyền vào GenerateToken
	userResponse := response.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role, // nếu user.Role là string hoặc user.Roles[0].Name thì chỉnh tương ứng
	}

	// 👇 Truyền thêm user.Role vào claim
	token, err := utils.GenerateToken(userResponse)
	if err != nil {
		return "", errors.New("Không thể tạo token")
	}

	return token, nil
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET")) // ⚠️ nhớ lấy từ biến môi trường
const RefreshableDuration = time.Hour * 24 * 7  // ví dụ 7 ngày

// Logout - vô hiệu hoá token (thêm vào DB)
func Logout(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("phương thức ký không hợp lệ")
		}
		return jwtSecret, nil
	})
	if err != nil {
		return exception.NewAppException(exception.Unauthenticated)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return exception.NewAppException(exception.Unauthenticated)
	}

	jti, ok := claims["jti"].(string)
	if !ok {
		return exception.NewAppException(exception.Unauthenticated)
	}

	expUnix, ok := claims["exp"].(float64)
	if !ok {
		return exception.NewAppException(exception.Unauthenticated)
	}

	expiryTime := time.Unix(int64(expUnix), 0)

	invalidToken := models.InvalidatedToken{
		ID:         jti,
		ExpiryTime: expiryTime,
	}

	if err := config.DB.Create(&invalidToken).Error; err != nil {
		return exception.NewAppException(exception.UncategorizedException)
	}

	return nil
}

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
	"github.com/jinzhu/copier"
)

func Login(req request.LoginRequest) (string, error) {
	var user models.User
	if err := config.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		return "", exception.NewAppException(exception.UserExisted)
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return "", errors.New("Mật khẩu không chính xác")
	}
	userResponse := response.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role, 
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


func Register(userRequest request.CreateUserRequest) (models.User, error) {
	// Check username tồn tại
	var existing models.User
	if err := config.DB.Where("username = ?", userRequest.Username).First(&existing).Error; err == nil {
		return models.User{}, exception.NewAppException(exception.UserExisted)
	}
	if len(userRequest.Username) < 3 {
		return models.User{}, exception.NewAppException(exception.UsernameInvalid)
	}
	if len(userRequest.Password) < 8 {
		return models.User{}, exception.NewAppException(exception.InvalidPassword)
	}
	if len(userRequest.FullName) < 2 || len(userRequest.FullName) > 50 {
		return models.User{}, exception.NewAppException(exception.InvalidFullname)
	}
	
	// Mã hoá mật khẩu
	hashedPassword, err := utils.HashPassword(userRequest.Password)
	if err != nil {
		return models.User{}, exception.NewAppException(exception.InvalidPassword)
	}

	var user models.User
	copier.Copy(&user, &userRequest)
	user.Password = hashedPassword

	if err := config.DB.Create(&user).Error; err != nil {
		return models.User{}, exception.NewAppException(exception.DatabaseError)
	}
	return user, nil
}
package services

import (
	"errors"

	"daoduy.com/hoc-golang/config"
	"daoduy.com/hoc-golang/models"
	request "daoduy.com/hoc-golang/models/request"
	"daoduy.com/hoc-golang/utils"
)

func Login(req request.LoginRequest) (string, error) {
	var user models.User
	if err := config.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		return "", errors.New("Người dùng không tồn tại")
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return "", errors.New("Mật khẩu không chính xác")
	}

	// 👇 Truyền thêm user.Role vào claim
	token, err := utils.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		return "", errors.New("Không thể tạo token")
	}

	return token, nil
}

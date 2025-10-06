package services

import (
	"errors"

	"daoduy.com/hoc-golang/config"
	"daoduy.com/hoc-golang/models"
	"daoduy.com/hoc-golang/models/request"
	"daoduy.com/hoc-golang/utils"
	"github.com/jinzhu/copier"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, errors.New("Không tìm thấy danh sách người dùng")
	}
	return users, nil
}

func CreateUser(userRequest request.CreateUserRequest) (models.User, error) {
	hashedPassword, err := utils.HashPassword(userRequest.Password)
	if err != nil {
		return models.User{}, errors.New("Không thể mã hoá mật khẩu")
	}
	var user models.User
	copier.Copy(&user, &userRequest)
	user.Password = hashedPassword

	if err := config.DB.Create(&user).Error; err != nil {
		return models.User{}, errors.New("Không thể tạo người dùng")
	}
	return user, nil
}

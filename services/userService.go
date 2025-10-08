package services

import (
	"daoduy.com/hoc-golang/config"
	"daoduy.com/hoc-golang/exception"
	"daoduy.com/hoc-golang/models"
	"daoduy.com/hoc-golang/models/request"
	"daoduy.com/hoc-golang/models/response"
	"daoduy.com/hoc-golang/utils"
	"github.com/jinzhu/copier"
)

func GetAllUsers() ([]response.UserResponse, error) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, exception.NewAppException(exception.UserNotFound)
	}

	var userResponses []response.UserResponse
	copier.Copy(&userResponses, &users)
	return userResponses, nil
}

func GetUserById(id string) (response.UserResponse, error) {
	var user models.User
	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return response.UserResponse{}, exception.NewAppException(exception.UserNotFound)
	}
	
	var userResponse response.UserResponse
	copier.Copy(&userResponse, &user)
	return userResponse, nil
}

func GetMyInfo(userID uint) (response.UserResponse, error) {
	var user models.User

	if err := config.DB.First(&user, userID).Error; err != nil {
		return response.UserResponse{}, exception.NewAppException(exception.UserNotFound)
	}

	var userResponse response.UserResponse
	copier.Copy(&userResponse, &user)
	// Map sang UserResponse
	return userResponse, nil
}

func CreateUser(userRequest request.CreateUserRequest) (models.User, error) {
	// Check username tồn tại
	var existing models.User
	if err := config.DB.Where("username = ?", userRequest.Username).First(&existing).Error; err == nil {
		return models.User{}, exception.NewAppException(exception.UserExisted)
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

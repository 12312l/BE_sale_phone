package controllers

import (
	"net/http"

	"daoduy.com/hoc-golang/models/request"
	"daoduy.com/hoc-golang/services"
	"github.com/gin-gonic/gin"
)

// @Summary Lấy danh sách tất cả người dùng
// @Description Trả về danh sách user trong hệ thống
// @Tags User
// @Produce json
// @Success 200 {array} models.User
// @Router /users/ [get]
func GetAllUsers(ctx *gin.Context) {
	users, err := services.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// @Summary Tạo mới người dùng
// @Description Tạo mới một người dùng trong hệ thống
// @Tags User
// @Accept json
// @Produce json
// @Param request body request.CreateUserRequest true "Thông tin tạo người dùng"
// @Success 200 {object} models.User
// @Failure 400 {string} string "Lỗi dữ liệu"
// @Router /users/ [post]
func CreateUser(ctx *gin.Context) {
	var user request.CreateUserRequest
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ"})
		return
	}

	newUser, err := services.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, newUser)
}

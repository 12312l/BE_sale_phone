package controllers

import (
	"net/http"

	"daoduy.com/hoc-golang/exception"
	"daoduy.com/hoc-golang/models/request"
	"daoduy.com/hoc-golang/models/response"
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
		ctx.Error(exception.NewAppException(exception.UncategorizedException))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(users))
}

// @Summary Lấy thông tin người dùng theo ID
// @Description Trả về thông tin chi tiết của người dùng theo ID
// @Tags User
// @Produce json
// @Param id path string true "ID của người dùng"
// @Success 200 {object} response.UserResponse
// @Router /users/{id} [get]
func GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := services.GetUserById(id)
	if err != nil {
		ctx.Error(exception.NewAppException(exception.UncategorizedException))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(user))
}

// @Summary Tạo mới người dùng
// @Description Tạo mới một người dùng trong hệ thống
// @Tags User
// @Accept json
// @Produce json
// @Param request body request.CreateUserRequest true "Thông tin tạo người dùng"
// @Success 201 {object} models.User
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Router /users/ [post]
func CreateUser(ctx *gin.Context) {
	var user request.CreateUserRequest
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.Error(exception.NewAppException(exception.InvalidKey))
		return
	}

	newUser, err := services.CreateUser(user)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, response.Success(newUser))
}

// MyInfoUser godoc
// @Summary Lấy thông tin người dùng hiện tại
// @Description Lấy thông tin người dùng từ token JWT
// @Tags User
// @Security BearerAuth
// @Produce json
// @Success 200 {object} response.UserResponse
// @Failure 404 {object} map[string]string
// @Router /users/myinfo [get]
func MyInfoUser(c *gin.Context) {
	userID := c.GetUint("user_id")


	user, err := services.GetMyInfo(userID)
	if err != nil {
		// c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		c.Error(exception.NewAppException(exception.Unauthenticated))
		return
	}

	c.JSON(http.StatusOK, response.Success(user))
}

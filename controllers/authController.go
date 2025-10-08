package controllers

import (
	"net/http"

	"daoduy.com/hoc-golang/exception"
	"daoduy.com/hoc-golang/models/request"
	"daoduy.com/hoc-golang/models/response"
	"daoduy.com/hoc-golang/services"
	"github.com/gin-gonic/gin"
)

// Login godoc
// @Summary Đăng nhập
// @Description Cho phép người dùng đăng nhập bằng username và password
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body request.LoginRequest true "Thông tin đăng nhập"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /auth/login [post]
func Login(ctx *gin.Context) {
	var req request.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(exception.NewAppException(exception.InvalidKey))
		return
	}

	token, err := services.Login(req)
	if err != nil {
		ctx.Error(exception.NewAppException(exception.Unauthorized))
		return
	}

	ctx.JSON(http.StatusOK, response.Success(gin.H{"token": token}))
}

// @Summary Đăng xuất người dùng
// @Description Thêm token vào danh sách vô hiệu hoá
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body request.LogoutRequest true "Token cần vô hiệu hoá"
// @Success 200 {object} response.ApiResponse
// @Router /auth/logout [post]
func Logout(ctx *gin.Context) {
	var req request.LogoutRequest

	// Lấy token từ body
	if err := ctx.ShouldBindJSON(&req); err != nil || req.Token == "" {
		ctx.Error(exception.NewAppException(exception.Unauthenticated))
		return
	}

	// Gọi service xử lý
	if err := services.Logout(req.Token); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response.Success(nil))
}

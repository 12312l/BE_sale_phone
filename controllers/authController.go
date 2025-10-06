package controllers

import (
	"net/http"

	"daoduy.com/hoc-golang/models/request"
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ"})
		return
	}

	token, err := services.Login(req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

package controllers

import (
	"net/http"

	"daoduy.com/hoc-golang/exception"
	"daoduy.com/hoc-golang/models/request"
	"daoduy.com/hoc-golang/models/response"
	"daoduy.com/hoc-golang/services"
	"github.com/gin-gonic/gin"
)

// @Summary Lấy danh sách tất cả danh mục
// @Description Lấy danh sách tất cả danh mục
// @Tags Category
// @Produce json
// @Success 200 {array} response.CategoryResponse
// @Router /categories/ [get]
func GetAllCategories(ctx *gin.Context) {
	categories, err := services.GetAllCategories()
	if err != nil {
		ctx.Error(exception.NewAppException(exception.UncategorizedException))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(categories))
}


// @Summary Tạo danh mục
// @Description Tạo danh mục
// @Tags Category
// @Accept json
// @Produce json
// @Param request body request.CategoryRequest true "Thông tin danh mục"
// @Success 201 {object} response.CategoryResponse
// @Router /categories/ [post]
func CreateCategory(ctx *gin.Context) {
	var category request.CategoryRequest
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.Error(exception.NewAppException(exception.InvalidKey))
		return
	}
	categoryResponse, err := services.CreateCategory(category)
	if err != nil {
		ctx.Error(exception.NewAppException(exception.UncategorizedException))
		return
	}
	ctx.JSON(http.StatusCreated, response.Success(categoryResponse))
}



// @Summary Lấy danh mục theo path
// @Description Lấy danh mục theo path
// @Tags Category
// @Produce json
// @Param path path string true "Path của danh mục"
// @Success 200 {object} response.CategoryResponse
// @Router /categories/{path} [get]
func GetCategoryByPath(ctx *gin.Context) {
	path := ctx.Param("path")
	categoryResponse, err := services.GetCategoryByPath(path)
	if err != nil {
		ctx.Error(exception.NewAppException(exception.UncategorizedException))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(categoryResponse))
}
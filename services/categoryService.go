package services

import (
	"strings"
	"unicode"

	"daoduy.com/hoc-golang/config"
	"daoduy.com/hoc-golang/exception"
	"daoduy.com/hoc-golang/models"
	"daoduy.com/hoc-golang/models/request"
	"daoduy.com/hoc-golang/models/response"
	"github.com/jinzhu/copier"
)

func GetAllCategories() ([]response.CategoryResponse, error) {
	var categories []models.Category
	if err := config.DB.Find(&categories).Error; err != nil {
		return nil, exception.NewAppException(exception.CategoryNotFound)
	}

	var categoryResponses []response.CategoryResponse
	copier.Copy(&categoryResponses, &categories)
	return categoryResponses, nil
}

func CreateCategory(categoryRequest request.CategoryRequest) (response.CategoryResponse, error) {
	var category models.Category
	copier.Copy(&category, &categoryRequest)
	category.Path = generateSlug(categoryRequest.Name)
	if err := config.DB.Create(&category).Error; err != nil {
		return response.CategoryResponse{}, exception.NewAppException(exception.DatabaseError)
	}

	var categoryResponse response.CategoryResponse
	copier.Copy(&categoryResponse, &category)
	return categoryResponse, nil
}

func GetCategoryByPath(path string) (response.CategoryResponse, error) {
	var category models.Category
	if err := config.DB.Where("path = ?", path).First(&category).Error; err != nil {
		return response.CategoryResponse{}, exception.NewAppException(exception.CategoryNotFound)
	}

	var categoryResponse response.CategoryResponse
	copier.Copy(&categoryResponse, &category)
	return categoryResponse, nil
}

func generateSlug(name string) string {
	// Mapping tiếng Việt có dấu thành không dấu
	vietnameseMap := map[rune]rune{
		'à': 'a', 'á': 'a', 'ạ': 'a', 'ả': 'a', 'ã': 'a', 'â': 'a', 'ầ': 'a', 'ấ': 'a', 'ậ': 'a', 'ẩ': 'a', 'ẫ': 'a', 'ă': 'a', 'ằ': 'a', 'ắ': 'a', 'ặ': 'a', 'ẳ': 'a', 'ẵ': 'a',
		'è': 'e', 'é': 'e', 'ẹ': 'e', 'ẻ': 'e', 'ẽ': 'e', 'ê': 'e', 'ề': 'e', 'ế': 'e', 'ệ': 'e', 'ể': 'e', 'ễ': 'e',
		'ì': 'i', 'í': 'i', 'ị': 'i', 'ỉ': 'i', 'ĩ': 'i',
		'ò': 'o', 'ó': 'o', 'ọ': 'o', 'ỏ': 'o', 'õ': 'o', 'ô': 'o', 'ồ': 'o', 'ố': 'o', 'ộ': 'o', 'ổ': 'o', 'ỗ': 'o', 'ơ': 'o', 'ờ': 'o', 'ớ': 'o', 'ợ': 'o', 'ở': 'o', 'ỡ': 'o',
		'ù': 'u', 'ú': 'u', 'ụ': 'u', 'ủ': 'u', 'ũ': 'u', 'ư': 'u', 'ừ': 'u', 'ứ': 'u', 'ự': 'u', 'ử': 'u', 'ữ': 'u',
		'ỳ': 'y', 'ý': 'y', 'ỵ': 'y', 'ỷ': 'y', 'ỹ': 'y',
		'đ': 'd',
		'À': 'a', 'Á': 'a', 'Ạ': 'a', 'Ả': 'a', 'Ã': 'a', 'Â': 'a', 'Ầ': 'a', 'Ấ': 'a', 'Ậ': 'a', 'Ẩ': 'a', 'Ẫ': 'a', 'Ă': 'a', 'Ằ': 'a', 'Ắ': 'a', 'Ặ': 'a', 'Ẳ': 'a', 'Ẵ': 'a',
		'È': 'e', 'É': 'e', 'Ẹ': 'e', 'Ẻ': 'e', 'Ẽ': 'e', 'Ê': 'e', 'Ề': 'e', 'Ế': 'e', 'Ệ': 'e', 'Ể': 'e', 'Ễ': 'e',
		'Ì': 'i', 'Í': 'i', 'Ị': 'i', 'Ỉ': 'i', 'Ĩ': 'i',
		'Ò': 'o', 'Ó': 'o', 'Ọ': 'o', 'Ỏ': 'o', 'Õ': 'o', 'Ô': 'o', 'Ồ': 'o', 'Ố': 'o', 'Ộ': 'o', 'Ổ': 'o', 'Ỗ': 'o', 'Ơ': 'o', 'Ờ': 'o', 'Ớ': 'o', 'Ợ': 'o', 'Ở': 'o', 'Ỡ': 'o',
		'Ù': 'u', 'Ú': 'u', 'Ụ': 'u', 'Ủ': 'u', 'Ũ': 'u', 'Ư': 'u', 'Ừ': 'u', 'Ứ': 'u', 'Ự': 'u', 'Ử': 'u', 'Ữ': 'u',
		'Ỳ': 'y', 'Ý': 'y', 'Ỵ': 'y', 'Ỷ': 'y', 'Ỹ': 'y',
		'Đ': 'd',
	}

	// Chuyển tiếng Việt có dấu thành không dấu
	slug := strings.Map(func(r rune) rune {
		if replacement, exists := vietnameseMap[r]; exists {
			return replacement
		}
		return r
	}, name)

	// Chuyển thành chữ thường
	slug = strings.ToLower(slug)

	// Thay thế khoảng trắng bằng dấu gạch ngang
	slug = strings.ReplaceAll(slug, " ", "-")

	// Xóa ký tự đặc biệt (chỉ giữ chữ và số, dấu gạch)
	slug = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '-' {
			return r
		}
		return -1
	}, slug)

	return slug
}

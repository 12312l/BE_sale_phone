package services

import (
	"daoduy.com/hoc-golang/config"
	"daoduy.com/hoc-golang/exception"
	"daoduy.com/hoc-golang/models"
	"daoduy.com/hoc-golang/models/response"
	"github.com/jinzhu/copier"
)

func GetAllProducts() ([]response.ProductResponse, error) {
	var products []models.Product
	if err := config.DB.Find(&products).Error; err != nil {
		return nil, exception.NewAppException(exception.ProductNotFound)
	}

	var productResponses []response.ProductResponse
	copier.Copy(&productResponses, &products)
	return productResponses, nil
}

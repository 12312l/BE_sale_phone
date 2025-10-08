package services

import (
	"daoduy.com/hoc-golang/config"
	"daoduy.com/hoc-golang/exception"
	"daoduy.com/hoc-golang/models"
	"daoduy.com/hoc-golang/models/request"
	"daoduy.com/hoc-golang/models/response"
	"github.com/jinzhu/copier"
)

func GetAddressByUserID(userID uint) ([]response.AddressResponse, error) {
	var addresses []models.Address

	if err := config.DB.Where("user_id = ?", userID).Find(&addresses).Error; err != nil {
		return nil, exception.NewAppException(exception.AddressNotFound)
	}

	if len(addresses) == 0 {
		return []response.AddressResponse{}, nil
	}

	var addressResponses []response.AddressResponse
	copier.Copy(&addressResponses, &addresses)

	return addressResponses, nil
}


func CreateAddressByUserID(userID uint, addressRequest request.AddressRequest) (response.AddressResponse, error) {
	var address models.Address
	copier.Copy(&address, &addressRequest)
	address.UserID = userID

	if err := config.DB.Create(&address).Error; err !=nil {
		return response.AddressResponse{}, exception.NewAppException(exception.DatabaseError)
	}

	var addressResponse response.AddressResponse
	copier.Copy(&addressResponse, &address)

	return addressResponse, nil
}

func UpdateAddressByUserID(addressID uint, userID uint, addressRequest request.AddressRequest) (response.AddressResponse, error) {
	var address models.Address

	if err := config.DB.Where("id = ? AND user_id = ?", addressID, userID).First(&address).Error; err != nil {
		return response.AddressResponse{}, exception.NewAppException(exception.AddressNotFound)
	}

	copier.Copy(&address, &addressRequest)

	if err := config.DB.Save(&address).Error; err != nil {
		return response.AddressResponse{}, exception.NewAppException(exception.DatabaseError)
	}

	var addressResponse response.AddressResponse
	copier.Copy(&addressResponse, &address)

	return addressResponse, nil
}

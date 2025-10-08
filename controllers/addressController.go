package controllers

import (
	"net/http"
	"strconv"

	"daoduy.com/hoc-golang/exception"
	"daoduy.com/hoc-golang/models/request"
	"daoduy.com/hoc-golang/models/response"
	"daoduy.com/hoc-golang/services"
	"github.com/gin-gonic/gin"
)

// @Summary      Lấy danh sách địa chỉ người dùng hiện tại
// @Description  Lấy danh sách tất cả địa chỉ của người dùng dựa trên token JWT
// @Tags         Address
// @Security     BearerAuth
// @Produce      json
// @Success      200 {array} response.AddressResponse
// @Failure      401 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Router       /address/myaddress [get]
func MyAddress(c *gin.Context) {
	userID := c.GetUint("user_id")

	address, err := services.GetAddressByUserID(userID)
	if err != nil {
		c.Error(exception.NewAppException(exception.Unauthenticated))
		return
	}
	c.JSON(http.StatusOK, response.Success(address))
}

// @Summary      Thêm địa chỉ mới cho người dùng hiện tại
// @Description  Người dùng có thể thêm địa chỉ mới dựa trên token JWT
// @Tags         Address
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        body  body  request.AddressRequest  true  "Thông tin địa chỉ"
// @Success      200 {object} response.AddressResponse
// @Failure      400 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /address/add [post]
func CreateAddress(c *gin.Context) {
	var addressRequest request.AddressRequest
	userID := c.GetUint("user_id")

	if err := c.ShouldBindJSON(&addressRequest); err != nil {
		c.Error(exception.NewAppException(exception.InvalidKey))
		return
	}

	newAddress, err := services.CreateAddressByUserID(userID, addressRequest)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.Success(newAddress))
}

// @Summary      Cập nhật địa chỉ hiện tại của người dùng
// @Description  Người dùng có thể cập nhật địa chỉ hiện tại dựa trên token JWT
// @Tags         Address
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "ID của địa chỉ"
// @Param        body  body  request.AddressRequest  true  "Thông tin địa chỉ"
// @Success      200 {object} response.AddressResponse
// @Failure      400 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /address/update/{id} [put]
func UpdateAddress(c *gin.Context) {
	var addressRequest request.AddressRequest
	userID := c.GetUint("user_id")
	addressIDStr := c.Param("id")

	addressID, err := strconv.ParseUint(addressIDStr, 10, 32)
	if err != nil {
		c.Error(exception.NewAppException(exception.InvalidID))
		return
	}

	if err := c.ShouldBindJSON(&addressRequest); err != nil {
		c.Error(exception.NewAppException(exception.InvalidKey))
		return
	}

	updatedAddress, err := services.UpdateAddressByUserID(uint(addressID), userID, addressRequest)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.Success(updatedAddress))
}

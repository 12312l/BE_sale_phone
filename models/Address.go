package models

type Address struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	FullName      string `json:"fullname"`
	Phone         string `json:"phone"`
	Province      string `json:"province"`
	District      string `json:"district"`
	Village       string `json:"village"`
	DetailAddress string `json:"detailaddress"`
	TypeAddress   bool   `json:"typeaddress"`

	// Khóa ngoại trỏ đến user
	UserID uint `json:"user_id"`
}

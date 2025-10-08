package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"fullname"`
	Phone    string `json:"phone"`
	Gender   string `json:"gender"`
	Dob      string `json:"dob"`
	Role     string `json:"role" gorm:"default:USER"`

	// Quan hệ 1-nhiều: 1 user có nhiều địa chỉ
	Addresses []Address `json:"addresses" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

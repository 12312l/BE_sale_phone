package models

import "time"

type User struct {
	ID       uint       `json:"id" gorm:"primaryKey"`
	Username string     `json:"username"`
	Email    string     `json:"email"`
	Password string     `json:"password"`
	FullName string     `json:"fullname"`
	Phone    string     `json:"phone"`
	Gender   string     `json:"gender"`
	Dob      *time.Time `json:"dob"`
	Role     string     `json:"role" gorm:"default:USER"`
}

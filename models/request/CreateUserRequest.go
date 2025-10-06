package request

import "time"

type CreateUserRequest struct {
	Username string     `json:"username" binding:"required"`
	Password string     `json:"password" binding:"required"`
	Email    string     `json:"email" binding:"required"`
	FullName string     `json:"fullname" binding:"required"`
	Phone    string     `json:"phone" binding:"required"`
	Gender   string     `json:"gender" binding:"required"`
	Dob      *time.Time `json:"dob" binding:"required"`
}

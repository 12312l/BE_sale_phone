package response

type UserResponse struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Email    string `json:"email"`
	FullName string `json:"fullname"`
	Phone    string `json:"phone"`
	Gender   string `json:"gender"`
	Dob      string `json:"dob"`
	Role     string `json:"role" gorm:"default:USER"`
}

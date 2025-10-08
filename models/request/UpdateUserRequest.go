package request

type UpdateUserRequest struct {
	FullName string `json:"fullname" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Gender   string `json:"gender" binding:"required"`
	Dob      string `json:"dob" binding:"required"`
}

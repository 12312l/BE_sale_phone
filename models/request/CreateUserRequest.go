package request

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
	FullName string `json:"fullname" binding:"required,min=2,max=50"`
	Phone    string `json:"phone" binding:"required"`
	Gender   string `json:"gender"`
	Dob      string `json:"dob"`
}

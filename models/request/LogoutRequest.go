package request

type LogoutRequest struct {
	Token string `json:"token" binding:"required"`
}

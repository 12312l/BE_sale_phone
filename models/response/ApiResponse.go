package response

type ApiResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Success tạo response thành công mặc định với Code = 1000
func Success(data interface{}) ApiResponse {
	return ApiResponse{
		Code: 1000,
		Data: data,
	}
}

// Error tạo response lỗi
func Error(code int, message string) ApiResponse {
	return ApiResponse{
		Code:    code,
		Message: message,
	}
}

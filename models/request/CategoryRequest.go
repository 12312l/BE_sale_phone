package request

// import "mime/multipart"

// type CategoryRequest struct {
// 	Name  string                `form:"name" binding:"required"`
// 	// Path  string                `form:"path" binding:"required"`
// 	Image *multipart.FileHeader `form:"image" binding:"required"`
// }
type CategoryRequest struct {
    Name     string `json:"name" binding:"required"`
	ImageURL string `json:"image_url" binding:"required"`
}


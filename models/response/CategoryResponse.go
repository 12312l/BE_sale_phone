package response

type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
	ImageURL string `json:"image_url"`
}

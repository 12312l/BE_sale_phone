package response

type SepecificationDetailResponse struct {
	ID   uint   `json:"id"`
	Content string `json:"content"`
	SepecificationID uint `json:"sepecification_id"`
}
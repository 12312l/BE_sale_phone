package response

type SepecificationResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	SepecificationDetails []SepecificationDetailResponse `json:"sepecification_details"`
}
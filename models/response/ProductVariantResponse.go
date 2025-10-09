package response

type ProductVariantResponse struct {
	ID   uint   `json:"id"`
	Type string `json:"type"`
	Price float64 `json:"price"`
	Stock int `json:"stock"`
	Images []ImageResponse `json:"images"`
	ColorID uint `json:"color_id"`
	Color ColorResponse `json:"color"`
	SepecificationID uint `json:"sepecification_id"`
	Sepecification SepecificationResponse `json:"sepecification"`
}
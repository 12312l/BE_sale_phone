package response

type ImageResponse struct {
	ID   uint   `json:"id"`
	URL  string `json:"url"`
	ProductVariantID uint `json:"product_variant_id"`
}
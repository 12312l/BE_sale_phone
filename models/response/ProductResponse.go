package response

type ProductResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	CategoryID uint `json:"category_id"`
	ProductVariants []ProductVariantResponse `json:"product_variants"`
}
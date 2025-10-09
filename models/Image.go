package models

type Image struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	URL  string `json:"url"`
	ProductVariantID uint `json:"product_variant_id"`
	ProductVariant ProductVariant `json:"product_variant" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

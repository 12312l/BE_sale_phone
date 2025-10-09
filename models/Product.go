package models

type Product struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Description string `json:"description"`
	CategoryID uint `json:"category_id"`
	ProductVariants []ProductVariant `json:"product_variants" gorm:"foreignKey:ProductID"`
	Category Category `json:"category" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
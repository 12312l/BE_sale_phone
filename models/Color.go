package models

type Color struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Code string `json:"code"`

	ProductVariants []ProductVariant `json:"product_variants" gorm:"foreignKey:ColorID"`
}

package models

type ProductVariant struct {
	ID               uint           `json:"id" gorm:"primaryKey"`
	Type             string         `json:"type"`
	Price            float64        `json:"price"`
	Stock            int            `json:"stock"`
	Images           []Image        `json:"images" gorm:"foreignKey:ProductVariantID"`
	ColorID          uint           `json:"color_id"`
	Color            Color          `json:"color" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SepecificationID uint           `json:"sepecification_id"`
	Sepecification   Sepecification `json:"sepecification" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductID        uint           `json:"product_id"`
	Product          Product        `json:"product" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

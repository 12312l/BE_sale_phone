package models

type Category struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	ImageURL string `json:"image_url"`

	Products []Product `json:"products" gorm:"foreignKey:CategoryID"`

}
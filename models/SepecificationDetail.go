package models

type SepecificationDetail struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Content string `json:"content"`
	SepecificationID uint `json:"sepecification_id"`
	Sepecification Sepecification `json:"sepecification" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
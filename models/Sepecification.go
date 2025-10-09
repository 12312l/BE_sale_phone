package models
type Sepecification struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`

	// 1–n với SepecificationDetail
	SepecificationDetails []SepecificationDetail `json:"sepecification_details" gorm:"foreignKey:SepecificationID"`
}

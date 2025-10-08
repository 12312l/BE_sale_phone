package models

import "time"

type InvalidatedToken struct {
	ID         string    `gorm:"primaryKey" json:"id"`
	ExpiryTime time.Time `json:"expiry_time"`
}

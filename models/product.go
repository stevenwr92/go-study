package models

import "time"

type Product struct {
	Id           uint   `gorm:"primaryKey" json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
	CreatedAt    time.Time
}

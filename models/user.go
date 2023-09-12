package models

import "time"

type User struct {
	Id        uint   `gorm:"primaryKey" json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	CreatedAt time.Time
}

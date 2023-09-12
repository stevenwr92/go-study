package models

import "time"

type Order struct {
	Id           uint    `gorm:"primaryKey" json:"id"`
	ProductRefer int     `json:"product_id"`
	Product      Product `gorm:"foreignKey:ProductRefer"`
	UserRefer    int     `json:"user_id"`
	User         User    `gorm:"foreignKey:UserRefer"`
	CreatedAt    time.Time
}

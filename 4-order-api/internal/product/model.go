package product

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null;uniqueIndex"`
	Description string `json:"description"`
	Price       int    `json:"price" gorm:"not null"`
	Quantity    int    `json:"quantity" gorm:"not null"`
	Image       string `json:"image"`
}

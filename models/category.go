package models

import (
	"gorm.io/gorm"
)

// Category represents a product category
type Category struct {
	gorm.Model
	Name     string    `json:"name"`
	Products []Product `json:"products" gorm:"foreignKey:CategoryID"`
}

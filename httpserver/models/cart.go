package models

import (
	"gorm.io/gorm"
)

// Cart represents a shopping cart
type Cart struct {
	gorm.Model
	UserID    uint       `json:"user_id"`
	CartItems []CartItem `json:"items" gorm:"foreignKey:CartID"`
	Total     float64    `json:"total"`
}

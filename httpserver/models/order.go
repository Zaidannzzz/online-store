package models

import (
	"gorm.io/gorm"
)

// Order represents an order made by a customer
type Order struct {
	gorm.Model
	UserID uint       `json:"user_id"`
	Total  float64    `json:"total"`
	Items  []CartItem `json:"items" gorm:"foreignKey:OrderID"`
	Status string     `json:"status"`
}

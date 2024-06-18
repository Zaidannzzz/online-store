package models

import (
	"gorm.io/gorm"
)

// CartItem represents an item in a shopping cart
type CartItem struct {
	gorm.Model
	CartID    uint    `json:"cart_id"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
	ProductID uint    `json:"product_id"`
	Quantity  uint    `json:"quantity"`
	Price     float64 `json:"price"`
	OrderID   uint    `json:"order_id"` // Foreign key for Order
}
